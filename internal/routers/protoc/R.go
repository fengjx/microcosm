package protoc

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
	return ErrMsg(RespCodeErr.msg)
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

func ResSuccess(c *gin.Context, data interface{}) {
	r := BuildData(RespCodeOk, data)
	c.JSON(http.StatusOK, r)
}

func ResError(c *gin.Context, msg string) {
	r := ErrMsg(msg)
	c.JSON(http.StatusInternalServerError, r)
}

func ResData(c *gin.Context, data R, httpcode int) {
	c.JSON(httpcode, data)
}

func ResCode(c *gin.Context, respCode RespCode, httpcode int) {
	r := Build(respCode)
	c.JSON(httpcode, r)
}

func Res(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Ok())
	} else {
		c.JSON(http.StatusInternalServerError, Err())
	}
}
