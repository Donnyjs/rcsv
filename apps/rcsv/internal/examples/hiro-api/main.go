package main

import (
	"fmt"
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/internal/config"
	"rcsv/apps/rcsv/internal/service/svc_inscription_monitor"
	"rcsv/domain/po"
	"rcsv/domain/repo"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/common/xredis"
	"rcsv/pkg/constant"
	"rcsv/pkg/utils"
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
	monitor := svc_inscription_monitor.NewInscriptionMonitor()
	monitor.FetchList(60, 0, constant.INSCRIPTION_LIST_ARGS)
}

func QueryInscriptionContent() {
	//92525    12720431   12720413  12597950
	monitor := svc_inscription_monitor.NewInscriptionMonitor()
	fmt.Println(monitor.Content("cdc1064c94785cd6617f5955f75a7939830ccec865fe85c0c91045efa6df5978i0"))
}

func InitData() {
	ticker := time.NewTicker(constant.INSCRIPTION_LIST_FETCH_INTERVAL)
	defer ticker.Stop()
	monitor := svc_inscription_monitor.NewInscriptionMonitor()
	repo := repo.NewInscriptionRepository()
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
			for _, v := range resp.Results {
				content, err := monitor.Content(v.Id)
				if err != nil {
					log.Errorf("query content err: %v, id: %s", err, v.Id)
					continue
				}
				//todo check svg data-clct="doodinals"
				_ = content
				var inscription po.Inscription
				inscription.Id = utils.NewUUID()
				inscription.Inscription = v.Number
				inscription.InscriptionId = v.Id
				err = repo.Insert(&inscription)
				if err != nil {
					log.Errorf("insert err: %v, id: %s, number: %d", err, v.Id, v.Number)
					continue
				}
			}
			offset += limit
			log.Infof("offset: %d", offset)
		}
	}
}
