package main

import (
	"fmt"
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/internal/config"
	"rcsv/apps/rcsv/internal/service/svc_inscription_monitor"
	"rcsv/domain/cache"
	"rcsv/domain/repo"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/common/xredis"
	"rcsv/pkg/constant"
	"rcsv/pkg/entity"
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
	QueryInscriptionContent()
	//InitData()
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
	content, _ := monitor.Content("6622fe6d83afec464193f6ab7155c74d5640e23d076a04d1e983e7581b112b81i0")
	fmt.Println(string(content))
	utils.ContainDataClctUtil(content)
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
			if resp.Total <= limit+offset && offset != 0 {
				log.Infof("init data done")
				return
			}
			throttle := make(chan struct{}, 5)
			var wg sync.WaitGroup
			log.Infof("1111")
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
					log.Infof("1111")
					flag, tp, list, _ := utils.ContainDataClctUtil(content)
					if !flag {
						return
					}
					_ = tp
					u := entity.NewMysqlUpdate()
					u.SetFilter("inscription=?", v.Number)
					u.Set("recursive_num", int64(len(list)))
					u.Set("genesis_timestamp", v.GenesisTimestamp)
					u.Set("genesis_block_height", v.GenesisBlockHeight)
					u.Set("content_length", v.ContentLength)
					err = repo.Update(u)
					if err != nil {
						log.Errorf("update err: %v, id: %s, number: %d", err, v.Id, v.Number)
						return
					}
					//cache.SetCurrentInscriptionNumber(int(v.Number))
				}(data)
			}
			wg.Wait()
			offset += limit
			log.Infof("offset: %d", offset)
		}
	}
}
