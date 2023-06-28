package svc_rcsv

import (
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/internal/dto/dto_collection"
	"rcsv/pkg/common/xhttp"
)

var log = logger.Logger("svc_collection")

func (s *collectionService) CollectionList(sort string, page, limit int) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)

	list, err := s.RcsvRepo.List(sort, page, limit)
	if err != nil {
		return nil
	}
	if err != nil {
		log.Error("list inscription err:", err)
		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, xhttp.ERROR_SERVER_INTERNAL_ERR)
		return
	}

	count, err := s.RcsvRepo.CountCollection()

	var dtoResp dto_collection.CollectionListResp
	dtoResp.DoodinalsTotal = count
	dtoResp.Data = list
	resp.Data = dtoResp
	log.Infof("resp: %v", resp)
	return

}
