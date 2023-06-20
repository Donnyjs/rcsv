package svc_inscription

import (
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/internal/dto/dto_inscription"
	"rcsv/pkg/common/xhttp"
	"rcsv/pkg/constant"
	"rcsv/pkg/entity"
)

var log = logger.Logger("svc_inscription")

func (s *inscriptionService) V2List(sort string, tp string, page int, limit int) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	//if s.InscriptionCache.ListExist() {
	//	list, err := s.InscriptionCache.GetInscriptionList(tp)
	//	if err != nil {
	//		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, err.Error())
	//		return resp
	//	}
	//	log.Infof("list: %v", list)
	//	var dtoResp []dto_inscription.InscriptionListResp
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

	w.SetSort(sort)
	w.SetOffset(limit * (page - 1))
	w.SetLimit(int32(limit))
	w.SetFilter("data_type=?", tp)
	list, err := s.InscriptionRepo.List(w)
	if err != nil {
		log.Error("list inscription err:", err)
		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, xhttp.ERROR_SERVER_INTERNAL_ERR)
		return
	}
	dCount, err := s.InscriptionRepo.CountDoodinals()
	if err != nil {
		log.Error("query count err:", err)
		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, xhttp.ERROR_SERVER_INTERNAL_ERR)
		return
	}
	oCount, err := s.InscriptionRepo.CountOthers()
	if err != nil {
		log.Error("query count err:", err)
		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, xhttp.ERROR_SERVER_INTERNAL_ERR)
		return
	}

	//s.InscriptionCache.SetInscriptionList(tp, list)
	var dtoResp dto_inscription.InscriptionListResp
	dtoResp.DoodinalsTotal = int(dCount)
	dtoResp.OthersTotal = int(oCount)

	for _, v := range list {
		dtoResp.Data = append(dtoResp.Data, dto_inscription.InscriptionData{
			Inscription:      int(v.Inscription),
			InscriptionId:    v.InscriptionId,
			ContentLength:    v.ContentLength,
			RecursiveNum:     v.RecursiveNum,
			GenesisTimestamp: v.GenesisTimestamp,
		})
	}
	resp.Data = dtoResp
	log.Infof("resp: %v", resp)
	return
}

func (s *inscriptionService) V1List(tp string, page int, limit int) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		w = entity.NewMysqlWhere()
	)

	w.SetSort(constant.INSCRIPTION_DESC)
	w.SetOffset(limit * (page - 1))
	w.SetLimit(int32(limit))
	w.SetFilter("data_type=?", tp)
	list, err := s.InscriptionRepo.List(w)
	if err != nil {
		log.Error("list inscription err:", err)
		resp.SetResult(xhttp.ERROR_CODE_SERVER_INTERNAL_ERR, xhttp.ERROR_SERVER_INTERNAL_ERR)
		return
	}
	//s.InscriptionCache.SetInscriptionList(tp, list)
	var dtoResp []dto_inscription.InscriptionData

	for _, v := range list {
		dtoResp = append(dtoResp, dto_inscription.InscriptionData{
			Inscription:   int(v.Inscription),
			InscriptionId: v.InscriptionId,
		})
	}
	resp.Data = dtoResp
	log.Infof("resp: %v", resp)
	return
}
