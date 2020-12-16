package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const SUCCESS = 0
const FAIL = 1

type Resp struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Detail interface{} `json:"detail,omitempty"`
}

func Response(ctx *gin.Context, StatusCode int, r Resp) {
	ctx.JSON(http.StatusOK, r)
	ctx.Abort()
}

func Success(ctx *gin.Context, msg string, data interface{}) {
	r := Resp{
		Code: SUCCESS,
		Msg:  msg,
		Data: data,
	}
	Response(ctx, http.StatusOK, r)
}

func Fail(ctx *gin.Context, bizCode int, msg string, data interface{}) {
	r := Resp{
		Code: bizCode,
		Msg:  msg,
		Data: data,
	}
	Response(ctx, http.StatusOK, r)
}

func FailWithErr(ctx *gin.Context, err error) {
	r := Resp{
		Code: FAIL,
		Msg:  "操作失败",
	}
	if cErr, ok := err.(*ErrCode); ok {
		r.Code = cErr.code
		r.Msg = GetErrInfo(cErr.code)
	}
	Response(ctx, http.StatusOK, r)
}
