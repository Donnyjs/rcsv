package svc_inscription_monitor

import (
	"rcsv/domain/cache"
	"rcsv/domain/po"
	"rcsv/domain/repo"
	"rcsv/pkg/constant"
	"rcsv/pkg/utils"
	"time"
)

type InscriptionMonitor struct {
	Cache cache.InscriptionCache
	Repo  repo.InscriptionRepository
}

func NewInscriptionMonitor(cache cache.InscriptionCache, repo repo.InscriptionRepository) *InscriptionMonitor {
	monitor := &InscriptionMonitor{
		Cache: cache,
		Repo:  repo,
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
			resp, err := im.FetchList(60, 0, constant.INSCRIPTION_LIST_ARGS)
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
