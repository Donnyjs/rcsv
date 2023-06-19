package svc_inscription

import (
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/internal/dto/dto_inscription"
	"rcsv/pkg/common/xhttp"
	"rcsv/pkg/entity"
)

var log = logger.Logger("svc_inscription")

func (s *inscriptionService) List(tp string, page int, limit int) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	//if s.InscriptionCache.ListExist() {
	//	list, err := s.InscriptionCache.GetInscriptionList()
	//	if err != nil {
	//		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, err.Error())
	//		return resp
	//	}
	//	dtoResp := make([]dto_inscription.InscriptionListResp, len(list))
	//	for _, v := range list {
	//		dtoResp = append(dtoResp, dto_inscription.InscriptionListResp{
	//			Inscription:   int(v.Inscription),
	//			InscriptionId: v.InscriptionId,
	//		})
	//	}
	//	log.Infof("data: %v, %d", dtoResp, len(list))
	//	resp.Data = dtoResp
	//	return
	//}

	var (
		w = entity.NewMysqlWhere()
	)

	w.SetSort("created_ts DESC")
	w.SetOffset(limit * (page - 1))
	w.SetLimit(int32(limit))
	w.SetFilter("data_type=?", tp)
	list, err := s.InscriptionRepo.List(w)
	if err != nil {
		log.Error("list inscription err:", err)
		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, xhttp.ERROR_SERVER_INTERNAL_ERR)
		return
	}
	//s.InscriptionCache.SetInscriptionList(list)
	var dtoResp []dto_inscription.InscriptionListResp
	for _, v := range list {
		dtoResp = append(dtoResp, dto_inscription.InscriptionListResp{
			Inscription:   int(v.Inscription),
			InscriptionId: v.InscriptionId,
		})
	}
	resp.Data = dtoResp
	log.Infof("resp: %v", resp)
	return
}
