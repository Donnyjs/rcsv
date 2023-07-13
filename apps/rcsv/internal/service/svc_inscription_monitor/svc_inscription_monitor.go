package svc_inscription_monitor

import (
	"fmt"
	"rcsv/domain/cache"
	"rcsv/domain/oss"
	"rcsv/domain/po"
	"rcsv/domain/repo"
	"rcsv/pkg/constant"
	"rcsv/pkg/utils"
	"time"
)

type InscriptionMonitor struct {
	Cache cache.InscriptionCache
	Repo  repo.InscriptionRepository
	Oss   oss.InscriptionOss
}

func NewInscriptionMonitor(cache cache.InscriptionCache, repo repo.InscriptionRepository, oss oss.InscriptionOss) *InscriptionMonitor {
	monitor := &InscriptionMonitor{
		Cache: cache,
		Repo:  repo,
		Oss:   oss,
	}
	return monitor
}

func (im *InscriptionMonitor) Run() {
	ticker := time.NewTicker(constant.INSCRIPTION_LIST_FETCH_INTERVAL)
	defer ticker.Stop()

	for {
		updateFlag := false
		select {
		case <-ticker.C:
			currentNumber := 0
			if im.Cache.NumberExist() {
				currentNumber = im.Cache.CurrentInscriptionNumber()
			}
			log.Infof("current number: %d", currentNumber)
			resp, err := im.FetchList(60, 0, fmt.Sprintf(constant.INSCRIPTION_LIST_NEW_ARGS, "%2B", currentNumber, 0))
			if err != nil {
				log.Errorf("fetchList err: %v", err)
				continue
			}
			for _, v := range resp.Results {
				if int(v.Number) <= currentNumber {
					continue
				}
				updateFlag = true
				content, err := im.Content(v.Id)
				if err != nil {
					log.Errorf("query content err: %v, id: %s", err, v.Id)
					continue
				}
				flag, tp, list, _ := utils.ContainDataClctUtil(content)
				if !flag {
					continue
				}
				var inscription po.Inscription
				inscription.Id = utils.NewUUID()
				inscription.Inscription = v.Number
				inscription.InscriptionId = v.Id
				inscription.DataType = tp
				inscription.ContentLength = v.ContentLength
				inscription.GenesisTimestamp = v.GenesisTimestamp
				inscription.GenesisBlockHeight = v.GenesisBlockHeight
				inscription.RecursiveNum = int64(len(list))
				inscription.Owner = v.Address
				picUrl, err := im.Oss.PutImage(&inscription)
				if err != nil {
					log.Error("putImage failure: ", err)
				}
				inscription.Pic = picUrl
				err = im.Repo.Insert(&inscription)
				if err != nil {
					log.Errorf("insert err: %v, id: %s, number: %d", err, v.Id, v.Number)
					continue
				}
				im.Cache.SetCurrentInscriptionNumber(int(v.Number))
			}
			if updateFlag {
				im.Cache.DeleteInscriptionList()
			}
		}
	}
}

func (im *InscriptionMonitor) RecursiveMonitor() {
	ticker := time.NewTicker(constant.INSCRIPTION_LIST_FETCH_INTERVAL)
	defer ticker.Stop()
	p_repo := repo.NewParentSonRepository()
	repoNew := repo.NewInscriptionNewRepository()

	for {
		select {
		case <-ticker.C:
			currentNumber := 0
			if im.Cache.RecursiveNumberExist() {
				currentNumber = im.Cache.CurrentRecursiveNumber()
			}
			log.Infof("current number: %d", currentNumber)
			resp, err := im.FetchList(60, 0, fmt.Sprintf(constant.FETCH_ALL_ARGS, "%2B", currentNumber, 0))
			if err != nil {
				log.Errorf("fetchList err: %v", err)
				continue
			}
			for _, v := range resp.Results {
				if int(v.Number) <= currentNumber {
					continue
				}
				var list []string
				content, err := im.Content(v.Id)
				if err != nil {
					log.Errorf("query content err: %v, id: %s", err, v.Id)
					continue
				}

				switch {
				case v.MimeType == constant.MIME_HTML:
					list, err = utils.ExcerptHTML(content)
				case v.MimeType == constant.MIME_SVG:
					_, _, list, err = utils.ContainDataClctUtil(content)
				}

				var inscription po.InscriptionNew
				inscription.Id = utils.NewUUID()
				inscription.Inscription = v.Number
				inscription.InscriptionId = v.Id
				inscription.ContentLength = v.ContentLength
				inscription.GenesisTimestamp = v.GenesisTimestamp
				inscription.GenesisBlockHeight = v.GenesisBlockHeight
				inscription.RecursiveNum = int64(len(list))
				inscription.Owner = v.Address

				err = repoNew.Insert(&inscription)
				if err != nil {
					log.Errorf("insert err: %v, id: %s, number: %d", err, v.Id, v.Number)
					continue
				}

				switch {
				case len(list) == 0:
					err = p_repo.Insert(&po.ParentSon{InsUuid: inscription.Id, InsId: v.Id, ParentInsId: v.Id})
					if err != nil {
						log.Errorf("insert parent_son err: %v, insuuid: %s, parentid: %s", err, inscription.Id, v.Id)
						continue
					}
				case len(list) > 0:
					var parentList []*po.ParentSon
					count := make(map[string]int, 0)
					for i := range list {
						if _, ok := count[list[i]]; ok {
							count[list[i]] += 1
						} else {
							count[list[i]] = 1
						}
					}

					for id, num := range count {
						parentList = append(parentList, &po.ParentSon{InsUuid: inscription.Id, InsId: id, ParentInsId: v.Id, Count: num})
					}

					err = p_repo.BatchInsert(parentList, len(parentList))
					if err != nil {
						log.Errorf("insert parent_son err: %v, insuuid: %s, parentid: %s", err, inscription.Id, v.Id)
						continue
					}
				}

				im.Cache.SetCurrentRecursiveNumber(int(v.Number))
			}
		}
	}
}
