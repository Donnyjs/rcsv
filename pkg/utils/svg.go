package utils

import (
	"encoding/xml"
	"errors"
	logger "github.com/ipfs/go-log"
	"rcsv/pkg/constant"
)

var log = logger.Logger("svg")

type InscriptionSvg struct {
	XMLName  string  `xml:"svg"`
	Style    string  `xml:"style,attr"`
	DataClct string  `xml:"data-clct,attr"`
	Version  string  `xml:"version,attr"`
	Image    []Image `xml:"g>image"`
}

type Image struct {
	XMLName xml.Name `xml:"image"`
	Href    string   `xml:"href,attr"`
}

func ContainDataClctUtil(body []byte) (bool, string, []string, error) {
	var (
		i  InscriptionSvg
		tp string
	)
	if err := xml.Unmarshal(body, &i); err != nil {
		log.Errorf("err: %v", err)
		return false, "", []string{}, err
	}

	if i.DataClct == constant.DATA_CLCT {
		log.Info("InscriptionInfo contains data-clct attribute with value 'doodinals'")
		tp = constant.DATA_CLCT
	}

	if i.DataClct == constant.DATA_RCSV_IO {
		log.Info("InscriptionInfo contains data-clct attribute with value 'rcsv.io'")
		tp = constant.DATA_RCSV_IO
	}
	if tp == "" {
		return false, "", []string{}, errors.New("InscriptionInfo not contains data-clct")
	}
	images := make([]string, 0)
	for index := range i.Image {
		images = append(images, i.Image[index].Href)
	}
	return true, tp, images, nil
}
