package ctrl_inscription

import (
	"github.com/gin-gonic/gin"
	"rcsv/pkg/common/xhttp"
)

func (ctrl *InscriptionCtrl) List(ctx *gin.Context) {
	var (
		resp *xhttp.Resp
	)
	resp = ctrl.inscriptionService.List()
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
