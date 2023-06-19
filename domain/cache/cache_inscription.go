package cache

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xredis"
	"rcsv/pkg/constant"
	"time"
)

type InscriptionCache interface {
	SetInscriptionList(resp []po.Inscription) (err error)
	GetInscriptionList() (resp []po.Inscription, err error)
	DeleteInscriptionList()
	ListExist() bool
	CurrentInscriptionNumber() int
	SetCurrentInscriptionNumber(number int)
	NumberExist() bool
}

type inscriptionCacheCache struct {
}

func NewInscriptionCache() InscriptionCache {
	return &inscriptionCacheCache{}
}

func (c *inscriptionCacheCache) SetInscriptionList(resp []po.Inscription) (err error) {
	var (
		key = constant.Inscription_List
	)
	log.Infof("resp: %v", resp)
	return Set(key, resp, time.Hour)
}

func (c *inscriptionCacheCache) GetInscriptionList() (resp []po.Inscription, err error) {
	var (
		key = constant.Inscription_List
	)
	resp = make([]po.Inscription, 10)
	err = Get(key, resp)
	return
}

func (c *inscriptionCacheCache) DeleteInscriptionList() {
	xredis.Del(constant.Inscription_List)
}

func (c *inscriptionCacheCache) ListExist() bool {
	return xredis.KeyExists(constant.Inscription_List)
}

func (c *inscriptionCacheCache) CurrentInscriptionNumber() int {
	number, _ := xredis.GetInt(constant.Inscription_NUMBER)
	return number
}

func (c *inscriptionCacheCache) SetCurrentInscriptionNumber(number int) {
	if number > c.CurrentInscriptionNumber() {
		Set(constant.Inscription_NUMBER, number, 0)
	}
}

func (c *inscriptionCacheCache) NumberExist() bool {
	return xredis.KeyExists(constant.Inscription_NUMBER)
}
