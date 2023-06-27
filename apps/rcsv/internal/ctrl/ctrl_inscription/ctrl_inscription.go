package ctrl_inscription

import (
	"rcsv/apps/rcsv/internal/service/svc_inscription"
	"rcsv/apps/rcsv/internal/service/svc_rcsv"
)

type InscriptionCtrl struct {
	inscriptionService svc_inscription.InscriptionService
	rcsvService        svc_rcsv.CollectionService
}

func NewInscriptionCtrl(inscriptionSvc svc_inscription.InscriptionService, rcsvService svc_rcsv.CollectionService) *InscriptionCtrl {
	return &InscriptionCtrl{inscriptionService: inscriptionSvc, rcsvService: rcsvService}
}
