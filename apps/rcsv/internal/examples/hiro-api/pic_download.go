package main

import (
	"io"
	"net/http"
	"os"
	"rcsv/domain/po"
	"rcsv/domain/repo"
	"strconv"
	"sync"
)

func DownloadPicSet(numStart int64, numEnd int64) {
	var wg sync.WaitGroup
	threshold := make(chan struct{}, 5)
	repo := repo.NewRCSVRepository()
	list, err := repo.ListWithNum(numStart, numEnd)
	if err != nil {
		log.Error(err)
	}
	stroagePath := "/Users/fuyiwei/Downloads/pic_set/"
	for _, item := range list {
		wg.Add(1)
		threshold <- struct{}{}
		go func(collection po.InscriptionInfoWithMetaNum) {

			// 发起HTTP GET请求获取图片数据
			resp, err := http.Get(collection.Pic)
			if err != nil {
				log.Error(err)
				return
			}

			// 创建本地文件
			file, err := os.Create(stroagePath + "Recursive Doodinal #" + strconv.FormatInt(collection.MetaNum, 10) + ".png")
			if err != nil {
				log.Error(err)
				return
			}

			// 将图片数据拷贝到本地文件
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				log.Error(err)
				return
			}

			defer func() {
				wg.Done()
				<-threshold
				resp.Body.Close()
				file.Close()
			}()
		}(item)
	}
	wg.Wait()
}
