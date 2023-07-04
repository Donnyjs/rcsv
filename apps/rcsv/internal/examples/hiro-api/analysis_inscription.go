package main

import (
	"fmt"
	"rcsv/apps/rcsv/internal/service/svc_inscription_monitor"
	"rcsv/domain/cache"
	oss2 "rcsv/domain/oss"
	"rcsv/domain/po"
	"rcsv/domain/repo"
	"rcsv/pkg/constant"
	"rcsv/pkg/utils"
	"sync"
	"time"
)

func AnalysisHistoryData() {
	ticker := make(chan struct{}, 1)
	p_repo := repo.NewParentSonRepository()
	repoNew := repo.NewInscriptionNewRepository()
	repo := repo.NewInscriptionRepository()
	cache := cache.NewInscriptionCache()
	oss := oss2.NewInscriptionOss()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo, oss)

	ticker <- struct{}{}
	var (
		limit      int64 = 60
		offset     int64 = 0
		fromNumber int64 = 1
		toNumber   int64 = 60000
	)
	for {
		select {
		case <-ticker:
			resp, err := monitor.FetchList(limit, offset, fmt.Sprintf(constant.FETCH_ALL_ARGS, "%2B", fromNumber, toNumber))
			log.Infof("query param total: %d, offset:%d, from: %d, to: %d", resp.Total, offset, fromNumber, toNumber)
			if err != nil {
				log.Errorf("fetchList err: %v, offset: %d", err, offset)
				continue
			}
			if resp.Total <= limit+offset && offset != 0 {
				log.Infof("init data done,total: %d, offset:%d, from: %d, to: %d", resp.Total, offset, fromNumber, toNumber)
				return
			}
			throttle := make(chan struct{}, 10)
			var wg sync.WaitGroup
			initTime := time.Now()
			for _, data := range resp.Results {
				wg.Add(1)
				throttle <- struct{}{}
				var list []string
				go func(v svc_inscription_monitor.Result) {
					defer func() {
						wg.Done()
						<-throttle
					}()

					content, err := monitor.Content(v.Id)
					if err != nil {
						log.Errorf("query content err: %v, id: %s", err, v.Id)
						return
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
						return
					}

					switch {
					case len(list) == 0:
						err = p_repo.Insert(&po.ParentSon{InsUuid: inscription.Id, InsId: v.Id, ParentInsId: v.Id})
						if err != nil {
							log.Errorf("insert parent_son err: %v, insuuid: %s, parentid: %s", err, inscription.Id, v.Id)
							return
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
							return
						}
					}

				}(data)
			}
			wg.Wait()
			offset += limit
			if resp.Total <= limit+offset && offset != 0 {
				log.Infof("End of current round, from: %d,to: %d", fromNumber, toNumber)
				fromNumber = toNumber + 1
				toNumber += 60000
				limit = 60
				offset = 0
			}
			ticker <- struct{}{}
			log.Infof("offset: %d,usedTime: %s", offset, time.Now().Sub(initTime).String())
		}
	}
}
