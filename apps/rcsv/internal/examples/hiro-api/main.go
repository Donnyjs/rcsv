package main

import (
	"fmt"
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/internal/config"
	"rcsv/apps/rcsv/internal/service/svc_inscription_monitor"
	"rcsv/domain/cache"
	"rcsv/domain/po"
	"rcsv/domain/repo"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/common/xredis"
	"rcsv/pkg/constant"
	"rcsv/pkg/utils"
	"sync"
	"time"
)

var log = logger.Logger("main")

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	logger.SetAllLoggers(logger.LevelInfo)
	//QueryInscriptionList()
	//QueryInscriptionContent()
	InitData()
}

func QueryInscriptionList() {
	//92525    12720431   12720413  12597950
	repo := repo.NewInscriptionRepository()
	cache := cache.NewInscriptionCache()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo)
	monitor.FetchList(60, 0, constant.INSCRIPTION_LIST_ARGS)
}

func QueryInscriptionContent() {
	//92525    12720431   12720413  12597950
	repo := repo.NewInscriptionRepository()
	cache := cache.NewInscriptionCache()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo)
	fmt.Println(monitor.Content("cdc1064c94785cd6617f5955f75a7939830ccec865fe85c0c91045efa6df5978i0"))
}

func InitData() {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	repo := repo.NewInscriptionRepository()
	cache := cache.NewInscriptionCache()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo)
	var (
		limit  int64 = 60
		offset int64 = 0
	)
	for {
		select {
		case <-ticker.C:
			resp, err := monitor.FetchList(limit, offset, constant.INSCRIPTION_LIST_ARGS+constant.INSCRIPTION_LIST_INIT_ARGS)
			if err != nil {
				log.Errorf("fetchList err: %v, offset: %d", err, offset)
				continue
			}
			if resp.Total <= limit+offset {
				log.Infof("init data done")
				return
			}
			throttle := make(chan struct{}, 5)
			var wg sync.WaitGroup
			for _, data := range resp.Results {
				wg.Add(1)
				throttle <- struct{}{}
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
					flag, tp, _ := utils.ContainDataClctUtil(content)
					if !flag {
						return
					}
					var inscription po.Inscription
					inscription.Id = utils.NewUUID()
					inscription.Inscription = v.Number
					inscription.InscriptionId = v.Id
					inscription.DataType = tp
					err = repo.Insert(&inscription)
					if err != nil {
						log.Errorf("insert err: %v, id: %s, number: %d", err, v.Id, v.Number)
						return
					}
					cache.SetCurrentInscriptionNumber(int(v.Number))
				}(data)
			}
			wg.Wait()
			offset += limit
			log.Infof("offset: %d", offset)
		}
	}
}
