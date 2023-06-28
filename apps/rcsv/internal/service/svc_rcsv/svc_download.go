package svc_rcsv

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"net/http"
	"rcsv/pkg/utils"
)

func (s *collectionService) DownloadPic(c *gin.Context, url string, width, height int64) {
	var imageBuf []byte
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx, utils.ScreenshotTasks(url, &imageBuf, width, height))
	if err != nil {
		log.Error("screenTask err:", err)
		return
	}

	c.Data(http.StatusOK, "image/png", imageBuf)
	return
}
