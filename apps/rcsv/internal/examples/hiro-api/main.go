package main

import (
	"encoding/json"
	"fmt"
	logger "github.com/ipfs/go-log"
	"github.com/tealeg/xlsx"
	"rcsv/apps/rcsv/internal/config"
	"rcsv/apps/rcsv/internal/service/svc_inscription_monitor"
	"rcsv/domain/cache"
	"rcsv/domain/repo"
	"rcsv/pkg/common/xmysql"
	"rcsv/pkg/common/xredis"
	"rcsv/pkg/constant"
	"rcsv/pkg/entity"
	"rcsv/pkg/utils"
	"strconv"
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
	//jiaban()
}

func jiaban() {
	repo := repo.NewInscriptionRepository()
	cache := cache.NewInscriptionCache()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo)
	_ = monitor
	data, _ := readFirstDataExcel()
	log.Infof("data: %v", data)
	var resp []Resp
	m := make(map[string]int64, 0)
	for _, v := range data {
		w := entity.NewMysqlWhere()
		w.SetFilter("inscription=?", v.Number)
		inscription, _ := repo.QueryByNumber(w)
		log.Infof("number: %d, id: %s", inscription.Inscription, inscription.InscriptionId)
		resp = append(resp, Resp{
			Id: inscription.InscriptionId,
			Meta: Metadata{
				Name: v.Name,
			},
		})
		//if inscription.Inscription == 0 {
		//	continue
		//}
		//time.Sleep(time.Second * 1)
		//content, _ := monitor.Content(inscription.InscriptionId)
		//if _, have := m[string(content)]; have {
		//	log.Infof("--------")
		//	log.Infof("id: %d", m[string(content)])
		//	m[string(content)] = inscription.Inscription
		//	log.Infof("id : %d", inscription.Inscription)
		//	log.Infof("--------")
		//} else {
		//	m[string(content)] = inscription.Inscription
		//}
	}

	log.Infof("len: %d", len(m))
	list, _ := json.Marshal(resp)
	log.Infof("list: %+v", string(list))
}

type Data struct {
	Number int
	Name   string
}

type Resp struct {
	Id   string   `json:"id"`
	Meta Metadata `json:"meta"`
}

type Metadata struct {
	Name string `json:"name"`
}

const path = "/Users/dongjs/Desktop/副本RCSV第一批200个入选.xlsx"

func readFirstDataExcel() ([]*Data, error) {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		log.Errorf("open file err : %v", err)
		return []*Data{}, err
	}
	var data []*Data
	for _, sheet := range xlFile.Sheets {
		log.Infof("sheet name : %s", sheet.Name)
		for _, row := range sheet.Rows {
			if row.Cells[1].String() == "number" {
				continue
			}
			number, _ := strconv.Atoi(row.Cells[1].String())

			log.Infof("number : %d, name : %s", number, row.Cells[2].String())
			data = append(data, &Data{
				Number: number,
				Name:   "Recursive Doodinal " + row.Cells[2].String(),
			})
		}
		log.Infof("\n")
	}
	return data, nil
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
	content, _ := monitor.Content("ba41b5324fd09681fb7a7be3ed8cca3154db0062d7cb1da09f723b34edae2688i0")
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
					flag, tp, list, _ := utils.ContainDataClctUtil(content)
					if !flag {
						return
					}
					_ = tp
					_ = list
					//var inscription po.Inscription
					//inscription.Id = utils.NewUUID()
					//inscription.Inscription = v.Number
					//inscription.InscriptionId = v.Id
					//inscription.DataType = tp
					//inscription.ContentLength = v.ContentLength
					//inscription.GenesisTimestamp = v.GenesisTimestamp
					//inscription.GenesisBlockHeight = v.GenesisBlockHeight
					//inscription.RecursiveNum = int64(len(list))
					//err = repo.Insert(&inscription)
					u := entity.NewMysqlUpdate()
					u.SetFilter("inscription=?", v.Number)
					u.Set("owner", v.Address)
					//u.Set("recursive_num", int64(len(list)))
					//u.Set("genesis_timestamp", v.GenesisTimestamp)
					//u.Set("genesis_block_height", v.GenesisBlockHeight)
					//u.Set("content_length", v.ContentLength)
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
