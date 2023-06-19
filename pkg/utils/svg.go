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

func ContainDataClctUtil(body []byte) (bool, []string, error) {
	var i InscriptionSvg
	if err := xml.Unmarshal(body, &i); err != nil {
		log.Error(err)
		return false, []string{}, err
	}

	if i.DataClct == constant.DATA_CLCT || i.DataClct == constant.RCSVIO {
		log.Info("InscriptionInfo contains data-clct attribute with value 'doodinals'")
		images := make([]string, 0)
		for index := range i.Image {
			images = append(images, i.Image[index].Href)
		}
		return true, images, nil
	}
	return false, []string{}, errors.New("InscriptionInfo not contains data-clct")
}
