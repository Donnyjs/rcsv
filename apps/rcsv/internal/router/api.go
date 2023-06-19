package router

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"rcsv/apps/rcsv/dig"
	"rcsv/apps/rcsv/internal/ctrl/ctrl_inscription"
	"rcsv/apps/rcsv/internal/service/svc_inscription"
)

var log = logger.Logger("router")

func Register(engine *gin.Engine) {
	publicGroup := engine.Group("api/v1")
	registerPublicRoutes(publicGroup)
}

func registerPublicRoutes(group *gin.RouterGroup) {
	registerInscriptionsRouter(group)
}

func registerInscriptionsRouter(group *gin.RouterGroup) {
	var svc svc_inscription.InscriptionService
	dig.Invoke(func(s svc_inscription.InscriptionService) {
		svc = s
	})
	ctrl := ctrl_inscription.NewInscriptionCtrl(svc)
	router := group.Group("inscription")
	router.GET("list", ctrl.List)
}
