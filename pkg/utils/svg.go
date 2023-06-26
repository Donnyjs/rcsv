package utils

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	logger "github.com/ipfs/go-log"
	loger "log"
	"rcsv/pkg/constant"
	"strconv"
	"time"
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

func ScreenShot(InscriptionId string) (imageBuf []byte, err error) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(loger.Printf))
	defer cancel()
	err = chromedp.Run(ctx, ScreenshotTasks(fmt.Sprintf(constant.INSCRIPTION_INFO, InscriptionId), &imageBuf))
	return imageBuf, err
}

func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EmulateViewport(2048, 2048),
		chromedp.Navigate(url),
		chromedp.FullScreenshot(imageBuf, 90),
	}
}

func NewObjectKey(inscriptionId int64) string {
	t := time.Now()
	timeFormat := t.Format("060102")
	return timeFormat + "/" + strconv.FormatInt(inscriptionId, 10) + ".png"
}
