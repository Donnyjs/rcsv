package utils

import (
	"encoding/xml"
	"errors"
	logger "github.com/ipfs/go-log"
	"rcsv/pkg/constant"
)

var log = logger.Logger("svg")

type InscriptionSvg struct {
	XMLName  string `xml:"svg"`
	Style    string `xml:"style,attr"`
	DataClct string `xml:"data-clct,attr"`
	Version  string `xml:"version,attr"`
}

func ContainDataClctUtil(body []byte) (bool, error) {
	var i InscriptionSvg
	if err := xml.Unmarshal(body, &i); err != nil {
		log.Error(err)
		return false, err
	}

	if i.DataClct == constant.DATA_CLCT {
		log.Info("InscriptionInfo contains data-clct attribute with value 'doodinals'")
		return true, nil
	}
	return false, errors.New("InscriptionInfo not contains data-clct")
}
