package ctrl_inscription

import (
	"github.com/gin-gonic/gin"
	"rcsv/apps/rcsv/internal/dto/dto_inscription"
	"rcsv/pkg/common/xhttp"
	"rcsv/pkg/constant"
)

func (ctrl *InscriptionCtrl) List(c *gin.Context) {
	var (
		resp *xhttp.Resp
	)
	var service dto_inscription.InscriptionListReq
	if err := c.BindQuery(&service); err == nil {
		if service.Type == constant.DOODINALS {

		} else if service.Type == constant.OTHERS {
			service.Type = constant.DATA_RCSV_IO
		} else {
			xhttp.Error(c, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR)
			return
		}
		resp = ctrl.inscriptionService.List(service.Type, service.Page, service.Limit)
		if resp.Code > 0 {
			xhttp.Error(c, resp.Code, resp.Msg)
			return
		}
		xhttp.Success(c, resp.Data)
	} else {
		xhttp.Error(c, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
	}
}
