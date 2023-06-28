package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"rcsv/apps/rcsv/internal/service/svc_inscription_monitor"
	"rcsv/domain/cache"
	oss2 "rcsv/domain/oss"
	"rcsv/domain/repo"
	"rcsv/pkg/constant"
	"rcsv/pkg/entity"
	"rcsv/pkg/utils"
	"strconv"
	"strings"
)

func ordyssey() {
	var sli = make([]map[string]interface{}, 700)
	err := json.Unmarshal([]byte(datajson), &sli)
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	log.Infof("resp: %v", sli)
	m := make(map[string]int64, 0)
	var idList []string
	for _, item := range sli {
		id, ok := item["id"].(string)
		idList = append(idList, id)
		if ok {
			if _, have := m[id]; have {
				log.Infof("--------")
				log.Infof("id: %s", id)
				log.Infof("--------")
			} else {
				m[id] = 1
			}
		}
	}

	log.Infof("idList: %d", len(idList))
	var resp []Resp
	var inscr []Inscription
	for k, v := range idList {
		k++
		resp = append(resp, Resp{
			Id: v,
			Meta: Metadata{
				Name: fmt.Sprintf("Recursive Doodinal #%d", k),
			},
		})
		inscr = append(inscr, Inscription{
			Id:     v,
			Number: fmt.Sprintf("Recursive Doodinal #%d", k),
		})
	}

	WriteDataExcel("/Users/dongjs/Desktop/1-654的副本.xlsx", inscr)
	list, _ := json.Marshal(resp)
	log.Infof("list: %+v", string(list))
}

func ordyssey1() {
	var sli = make([]map[string]interface{}, 700)
	err := json.Unmarshal([]byte(datajson), &sli)
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	log.Infof("resp: %v", sli)
	m := make(map[string]int64, 0)
	var idList []string
	for _, item := range sli {
		id, ok := item["id"].(string)
		idList = append(idList, id)
		if ok {
			if _, have := m[id]; have {
				log.Infof("--------")
				log.Infof("id: %s", id)
				log.Infof("--------")
			} else {
				m[id] = 1
			}
		}
	}

	log.Infof("idList: %d", len(idList))
	for k, v := range idList {
		k++
		fmt.Printf("%s, Genesis Punk #%d", v, k)
		fmt.Println()
	}

}

func jiaban() {
	repo := repo.NewInscriptionRepository()
	cache := cache.NewInscriptionCache()
	oss := oss2.NewInscriptionOss()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo, oss)
	_ = monitor
	data, _ := readFirstDataExcel()
	log.Infof("data: %v", data)
	var resp []Resp
	//m := make(map[string]int64, 0)
	d := make(map[string]string, 0)
	for _, v := range data {
		w := entity.NewMysqlWhere()
		w.SetFilter("inscription_id=?", v.Id)
		inscription, _ := repo.QueryByNumber(w)
		log.Infof("number: %d, id: %s", inscription.Inscription, inscription.InscriptionId)
		resp = append(resp, Resp{
			Id: inscription.InscriptionId,
			Meta: Metadata{
				Name: v.Name,
			},
		})
		if inscription.Inscription == 0 {
			continue
		}
		//content, _ := monitor.Content(inscription.InscriptionId)
		//flag, _, _, _ := utils.ContainDataClctUtil(content)
		//if !flag {
		//	log.Errorf("err: ======= %v", inscription.InscriptionId)
		//	return
		//}
		//
		//if _, have := m[string(content)]; have {
		//	log.Infof("--------")
		//	log.Infof("id: %d", m[string(content)])
		//	m[string(content)] = inscription.Inscription
		//	log.Infof("id : %d", inscription.Inscription)
		//	log.Infof("--------")
		//} else {
		//	m[string(content)] = inscription.Inscription
		//}

		if _, have := d[inscription.Owner]; have {
			log.Infof("========")
			id := d[inscription.Owner]
			id = id + "," + inscription.InscriptionId
			d[inscription.Owner] = id
			log.Infof("=======")
		} else {
			d[inscription.Owner] = inscription.InscriptionId
		}
	}

	list, _ := json.Marshal(resp)
	WritWteDataExcel("/Users/dongjs/Desktop/问题说明.xlsx", d)
	log.Infof("list: %+v", string(list))
}

type Data struct {
	Number int
	Id     string
	Name   string
}

type Wenti struct {
	Owner string
	Id    string
}

type R struct {
	R []Resp `json:"resp"`
}

type Resp struct {
	Id   string   `json:"id"`
	Meta Metadata `json:"meta"`
}

type Metadata struct {
	Name string `json:"name"`
}

const path = "/Users/dongjs/Desktop/1-641.xlsx"

func readFirstDataExcel() ([]*Data, error) {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		log.Errorf("open file err : %v", err)
		return []*Data{}, err
	}
	var data []*Data
	name := 0
	for _, sheet := range xlFile.Sheets {
		log.Infof("sheet name : %s", sheet.Name)
		for _, row := range sheet.Rows {
			if row.Cells[0].String() == "id" {
				continue
			}
			name++
			id := row.Cells[0].String()

			log.Infof("id : %s", id)
			data = append(data, &Data{
				Id:   id,
				Name: "Recursive Doodinal #" + strconv.Itoa(name),
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
	oss := oss2.NewInscriptionOss()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo, oss)
	monitor.FetchList(60, 0, constant.INSCRIPTION_LIST_ARGS)
}

func QueryInscriptionContent() {
	repo := repo.NewInscriptionRepository()
	cache := cache.NewInscriptionCache()
	oss := oss2.NewInscriptionOss()
	monitor := svc_inscription_monitor.NewInscriptionMonitor(cache, repo, oss)
	content, _ := monitor.Content("21475eb3a0512d5cf9812e0ca129e47bf4e12fe681fad3cc28de0f62349e38a1i0")
	fmt.Println(string(content))
	utils.ContainDataClctUtil(content)
}

type Inscription struct {
	Id     string
	Number string
}

func WriteDataExcel(path string, data []Inscription) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("remove path : ", err)
		return
	}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("data")
	if err != nil {
		fmt.Println("add sheet", err)
		return
	}
	row := sheet.AddRow()

	cell := row.AddCell()
	cell.Value = "id"
	cell = row.AddCell()
	cell.Value = "number"

	for _, info := range data {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = info.Id
		cell = row.AddCell()
		cell.Value = info.Number
	}
	err = file.Save(path)
	if err != nil {
		panic(err.Error())
	}
}

func WritWteDataExcel(path string, m map[string]string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("remove path : ", err)
		return
	}
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("data")
	if err != nil {
		fmt.Println("add sheet", err)
		return
	}
	row := sheet.AddRow()

	cell := row.AddCell()
	cell.Value = "owner"
	cell = row.AddCell()
	cell.Value = "id"

	for k, v := range m {
		if strings.Contains(v, ",") {
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = k
			cell = row.AddCell()
			cell.Value = v
		}
	}
	err = file.Save(path)
	if err != nil {
		panic(err.Error())
	}
}
