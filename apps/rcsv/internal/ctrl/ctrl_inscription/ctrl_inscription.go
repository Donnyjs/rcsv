package ctrl_inscription

import (
	"rcsv/apps/rcsv/internal/service/svc_inscription"
)

type InscriptionCtrl struct {
	inscriptionService svc_inscription.InscriptionService
}

func NewInscriptionCtrl(inscriptionSvc svc_inscription.InscriptionService) *InscriptionCtrl {
	return &InscriptionCtrl{inscriptionService: inscriptionSvc}
}
