package svc_inscription_monitor

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"rcsv/pkg/constant"
)

func (im *InscriptionMonitor) Content(id string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "image/svg+xml").
		Get(fmt.Sprintf(constant.INSCRIPTION_INFO, id))
	if err != nil {
		log.Errorf("err : %v", err)
		return "", err
	}
	return string(resp.Body()), nil
}
