package svc_rcsv

import (
	"github.com/gin-gonic/gin"
	"rcsv/domain/repo"
	"rcsv/pkg/common/xhttp"
)

type CollectionService interface {
	CollectionList(sort string, page, limit int) (resp *xhttp.Resp)
	DownloadPic(c *gin.Context, url string, width, height int64)
}

type collectionService struct {
	RcsvRepo repo.RCSVRepository
}

func NewCollectionService(repository repo.RCSVRepository) CollectionService {
	return &collectionService{
		RcsvRepo: repository,
	}
}
