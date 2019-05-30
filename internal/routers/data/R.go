package data

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok() *R {
	return OkMsg(RespCodeOk.msg)
}

func OkMsg(msg string) *R {
	r := new(R)
	r.Code = RespCodeOk.code
	r.Msg = msg
	return r
}

func Err() *R {
	return OkMsg(RespCodeErr.msg)
}

func ErrMsg(msg string) *R {
	r := new(R)
	r.Code = RespCodeErr.code
	r.Msg = msg
	return r
}

func Build(respCode RespCode) *R {
	r := new(R)
	r.Code = respCode.code
	r.Msg = respCode.msg
	return r
}

func BuildData(respCode RespCode, data interface{}) *R {
	r := new(R)
	r.Code = respCode.code
	r.Msg = respCode.msg
	r.Data = data
	return r
}

func Res(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Ok())
	} else {
		c.JSON(http.StatusInternalServerError, Ok())
	}
}
