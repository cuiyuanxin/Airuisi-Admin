package app

import (
	"github.com/cuiyuanxin/airuisi-admin/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

//type Pager struct {
//	Page      int `json:"pageNo"`
//	Limit     int `json:"pageSize"`
//	TotalRows int `json:"totalCount"`
//	TotalPage int `json:"totalPage"`
//}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToSuccess(data interface{}) {
	resp := gin.H{"code": 0, "msg": "success"}
	if data != nil {
		resp["data"] = data
	}

	r.Ctx.JSON(http.StatusOK, resp)
}

func (r *Response) ToError(data interface{}, err *errcode.Error) {
	resp := gin.H{"code": err.Code(), "msg": err.Msg()}
	if data != nil {
		resp["data"] = data
	}

	r.Ctx.JSON(err.StatusCode(), resp)
}

func (r *Response) ToResponse(data interface{}, err *errcode.Error) {
	resp := gin.H{"code": err.Code(), "msg": err.Msg()}
	if data != nil {
		resp["data"] = data
	}
	r.Ctx.JSON(http.StatusOK, resp)
}

//func (r *Response) ToResponseList(list interface{}, totalRows, totalPage int) {
//	r.Ctx.JSON(http.StatusOK, gin.H{
//		"list": list,
//		"pager": Pager{
//			Page:      GetPage(r.Ctx),
//			Limit:     GetLimit(r.Ctx),
//			TotalRows: totalRows,
//			TotalPage: totalPage,
//		},
//	})
//}
