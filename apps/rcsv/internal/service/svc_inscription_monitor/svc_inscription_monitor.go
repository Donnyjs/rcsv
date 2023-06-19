package svc_inscription_monitor

import (
	"rcsv/pkg/constant"
	"time"
)

type InscriptionMonitor struct {
}

func NewInscriptionMonitor() *InscriptionMonitor {
	ticker := time.NewTicker(constant.INSCRIPTION_LIST_FETCH_INTERVAL)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Infof("11111")
			}
		}
	}()
	return &InscriptionMonitor{}
}
