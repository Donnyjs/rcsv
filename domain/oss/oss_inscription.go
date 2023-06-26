package oss

import (
	"rcsv/apps/rcsv/config"
	"rcsv/domain/po"
	"rcsv/pkg/common/xoss"
	"rcsv/pkg/utils"
	"strings"
)

type InscriptionOss interface {
	PutImage(inscription *po.Inscription) (fileName string, err error)
}

type inscriptionOss struct {
}

func NewInscriptionOss() InscriptionOss {
	return &inscriptionOss{}
}

func (i *inscriptionOss) PutImage(inscription *po.Inscription) (fileName string, err error) {

	imageBuf, err := utils.ScreenShot(inscription.InscriptionId)
	fileName = utils.NewObjectKey(inscription.Inscription)
	err = xoss.GetOss().PutObject(fileName, strings.NewReader(string(imageBuf)))
	if err != nil {
		return "", err
	}

	return config.GetConfig().Oss.Domain + fileName, nil
}
