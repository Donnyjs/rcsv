package svc_inscription

import (
	logger "github.com/ipfs/go-log"
	"rcsv/pkg/common/xhttp"
	"rcsv/pkg/entity"
)

var log = logger.Logger("svc_inscription")

func (s *inscriptionService) List() (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		w = entity.NewMysqlWhere()
	)
	w.SetFilter("1=?", 1)
	w.SetSort("created_ts DESC")
	w.SetLimit(3)
	list, err := s.InscriptionRepo.List(w)
	if err != nil {
		log.Error("list inscription err:", err)
		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, xhttp.ERROR_SERVER_INTERNAL_ERR)
		return
	}
	resp.Data = list
	log.Infof("resp: %v", resp)
	return
}
