package cache

import (
	"rcsv/domain/po"
	"rcsv/pkg/constant"
)

type InscriptionCache interface {
	SetInscriptionList(resp []po.Inscription) (err error)
	GetInscriptionList() (resp []po.Inscription, err error)
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
	return Set(key, resp, 0)
}

func (c *inscriptionCacheCache) GetInscriptionList() (resp []po.Inscription, err error) {
	var (
		key = constant.Inscription_List
	)
	resp = make([]po.Inscription, 0)
	err = Get(key, resp)
	return
}
