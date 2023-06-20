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
	v1PublicGroup := engine.Group("api/v1")
	registerV1PublicRoutes(v1PublicGroup)

	v2PublicGroup := engine.Group("api/v2")
	registerV2PublicRoutes(v2PublicGroup)
}

func registerV1PublicRoutes(group *gin.RouterGroup) {
	registerV1InscriptionsRouter(group)
}

func registerV2PublicRoutes(group *gin.RouterGroup) {
	registerV2InscriptionsRouter(group)
}

func registerV1InscriptionsRouter(group *gin.RouterGroup) {
	var svc svc_inscription.InscriptionService
	dig.Invoke(func(s svc_inscription.InscriptionService) {
		svc = s
	})
	ctrl := ctrl_inscription.NewInscriptionCtrl(svc)
	router := group.Group("inscription")
	router.GET("list", ctrl.V1List)
}

func registerV2InscriptionsRouter(group *gin.RouterGroup) {
	var svc svc_inscription.InscriptionService
	dig.Invoke(func(s svc_inscription.InscriptionService) {
		svc = s
	})
	ctrl := ctrl_inscription.NewInscriptionCtrl(svc)
	router := group.Group("inscription")
	router.GET("list", ctrl.V2List)
}
