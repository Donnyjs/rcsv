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

func ContainDataClctUtil(body []byte) (bool, string, error) {
	var i InscriptionSvg
	if err := xml.Unmarshal(body, &i); err != nil {
		return false, "", err
	}

	if i.DataClct == constant.DATA_CLCT {
		log.Info("InscriptionInfo contains data-clct attribute with value 'doodinals'")
		images := make([]string, 0)
		for index := range i.Image {
			images = append(images, i.Image[index].Href)
		}
		_ = images
		return true, constant.DATA_CLCT, nil
	}

	if i.DataClct == constant.DATA_RCSV_IO {
		images := make([]string, 0)
		for index := range i.Image {
			images = append(images, i.Image[index].Href)
		}
		_ = images
		log.Info("InscriptionInfo contains data-clct attribute with value 'rcsv.io'")
		return true, constant.DATA_RCSV_IO, nil
	}

	return false, "", errors.New("InscriptionInfo not contains data-clct")
}
