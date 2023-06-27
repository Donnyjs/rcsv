package ctrl_inscription

import (
	"github.com/gin-gonic/gin"
	"rcsv/apps/rcsv/internal/dto/dto_inscription"
	"rcsv/pkg/common/xhttp"
	"rcsv/pkg/constant"
)

func (ctrl *InscriptionCtrl) V1List(c *gin.Context) {
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

		resp = ctrl.inscriptionService.V1List(service.Type, service.Page, service.Limit)
		if resp.Code > 0 {
			xhttp.Error(c, resp.Code, resp.Msg)
			return
		}
		xhttp.Success(c, resp.Data)
	} else {
		xhttp.Error(c, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
	}
}

func (ctrl *InscriptionCtrl) V2List(c *gin.Context) {
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
		if service.Sort == constant.LEAST_RECURSIONS {
			service.Sort = constant.RECURSIONS_ASC
		} else if service.Sort == constant.Most_RECURSIONS {
			service.Sort = constant.RECURSIONS_DESC
		} else if service.Sort == constant.NEWEST {
			service.Sort = constant.INSCRIPTION_DESC
		} else if service.Sort == constant.OLDEST {
			service.Sort = constant.INSCRIPTION_ASC
		} else {
			xhttp.Error(c, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR)
			return
		}
		resp = ctrl.inscriptionService.V2List(service.Sort, service.Type, service.Page, service.Limit)
		if resp.Code > 0 {
			xhttp.Error(c, resp.Code, resp.Msg)
			return
		}
		xhttp.Success(c, resp.Data)
	} else {
		xhttp.Error(c, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
	}
}

func (ctrl *InscriptionCtrl) CollectionList(c *gin.Context) {
	var (
		resp *xhttp.Resp
	)
	var service dto_inscription.InscriptionListReq
	if err := c.BindQuery(&service); err == nil {
		if service.Sort == constant.LEAST_RECURSIONS {
			service.Sort = constant.RECURSIONS_TABLE_ASC
		} else if service.Sort == constant.Most_RECURSIONS {
			service.Sort = constant.RECURSIONS_TABLE_DESC
		} else if service.Sort == constant.NEWEST {
			service.Sort = constant.INSCRIPTION_TABLE_DESC
		} else if service.Sort == constant.OLDEST {
			service.Sort = constant.INSCRIPTION_TABLE_ASC
		} else if service.Sort == constant.DEFAULT_NEWEST {
			service.Sort = constant.METANUMD_DESC
		} else if service.Sort == constant.DEFAULT_OLDEST {
			service.Sort = constant.METANUMD_ASC
		} else {
			xhttp.Error(c, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR)
			return
		}
		resp = ctrl.rcsvService.CollectionList(service.Sort, service.Page, service.Limit)
		if resp.Code > 0 {
			xhttp.Error(c, resp.Code, resp.Msg)
			return
		}
		xhttp.Success(c, resp.Data)
	} else {
		xhttp.Error(c, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
	}
}
