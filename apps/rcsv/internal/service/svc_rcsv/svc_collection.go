package svc_rcsv

import (
	"rcsv/domain/repo"
	"rcsv/pkg/common/xhttp"
)

type CollectionService interface {
	CollectionList(sort string, page, limit int) (resp *xhttp.Resp)
}

type collectionService struct {
	RcsvRepo repo.RCSVRepository
}

func NewCollectionService(repository repo.RCSVRepository) CollectionService {
	return &collectionService{
		RcsvRepo: repository,
	}
}
