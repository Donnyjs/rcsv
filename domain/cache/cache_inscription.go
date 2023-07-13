package cache

import (
	"rcsv/domain/po"
	"rcsv/pkg/common/xredis"
	"rcsv/pkg/constant"
	"rcsv/pkg/utils"
)

type InscriptionCache interface {
	SetInscriptionList(tp string, resp []po.Inscription) (err error)
	GetInscriptionList(tp string) (resp []po.Inscription, err error)
	DeleteInscriptionList()
	ListExist() bool
	CurrentInscriptionNumber() int
	SetCurrentInscriptionNumber(number int)
	NumberExist() bool
	CurrentRecursiveNumber() int
	SetCurrentRecursiveNumber(number int)
	RecursiveNumberExist() bool
}

type inscriptionCacheCache struct {
}

func NewInscriptionCache() InscriptionCache {
	return &inscriptionCacheCache{}
}

func (c *inscriptionCacheCache) SetInscriptionList(tp string, resp []po.Inscription) (err error) {
	var (
		key = constant.Inscription_List + tp
	)
	log.Infof("resp: %v", resp)
	return Set(key, resp, 0)
}

func (c *inscriptionCacheCache) GetInscriptionList(tp string) (resp []po.Inscription, err error) {
	var (
		key = constant.Inscription_List + tp
	)
	var list []po.Inscription
	jsonStr, err := xredis.Get(key)
	log.Infof("str: %s ", jsonStr)
	if err != nil {
		log.Error(ERROR_CACHE_REDIS_GET_FAILED, err.Error())
		return
	}
	if jsonStr == "" {
		return
	}
	err = utils.Unmarshal(jsonStr, &list)
	log.Infof("out: %v", list)
	if err != nil {
		log.Warn(ERROR_CACHE_PROTOCOL_UNMARSHAL_ERR, err.Error())
	}
	return list, nil
}

func (c *inscriptionCacheCache) DeleteInscriptionList() {
	xredis.DelByPrefix(constant.Inscription_List)
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

func (c *inscriptionCacheCache) CurrentRecursiveNumber() int {
	number, _ := xredis.GetInt(constant.RECURSIVE_NUMBER)
	return number
}

func (c *inscriptionCacheCache) SetCurrentRecursiveNumber(number int) {
	if number > c.CurrentRecursiveNumber() {
		Set(constant.RECURSIVE_NUMBER, number, 0)
	}
}

func (c *inscriptionCacheCache) RecursiveNumberExist() bool {
	return xredis.KeyExists(constant.RECURSIVE_NUMBER)
}
