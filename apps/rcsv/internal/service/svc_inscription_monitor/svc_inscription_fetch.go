package svc_inscription_monitor

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	logger "github.com/ipfs/go-log"
	"rcsv/pkg/constant"
)

var log = logger.Logger("svc_inscription_monitor")

func (im *InscriptionMonitor) FetchList(limit, offset int64, args string) (hir HiroInscriptionResp, err error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf(constant.INSCRIPTION_LIST, limit, offset, args))
	if err != nil {
		log.Errorf("err : %v", err)
		return
	}
	err = json.Unmarshal(resp.Body(), &hir)
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	log.Infof("hir : %v", hir)
	return hir, nil
}
