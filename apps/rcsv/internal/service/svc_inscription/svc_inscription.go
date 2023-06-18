package svc_inscription

import (
	"rcsv/domain/cache"
	"rcsv/domain/repo"
	"rcsv/pkg/common/xhttp"
)

type InscriptionService interface {
	List() (resp *xhttp.Resp)
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
