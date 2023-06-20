package svc_inscription

import (
	"rcsv/domain/cache"
	"rcsv/domain/repo"
	"rcsv/pkg/common/xhttp"
)

type InscriptionService interface {
	V1List(tp string, page, limit int) (resp *xhttp.Resp)
	V2List(sort, tp string, page, limit int) (resp *xhttp.Resp)
}

type inscriptionService struct {
	InscriptionCache cache.InscriptionCache
	InscriptionRepo  repo.InscriptionRepository
}

func NewInscriptionService(inscriptionCache cache.InscriptionCache, repository repo.InscriptionRepository) InscriptionService {
	return &inscriptionService{
		InscriptionCache: inscriptionCache,
		InscriptionRepo:  repository,
	}
}
