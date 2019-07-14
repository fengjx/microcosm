package protoc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// R 对象定义数据返回的解构，统一接口返回格式
type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Ok 构造成功消息，只有响应码，没有message
func Ok() *R {
	return OkMsg(RespCodeOk.msg)
}

// OkMsg 构造成功消息，可以自定义返回消息
func OkMsg(msg string) *R {
	r := new(R)
	r.Code = RespCodeOk.code
	r.Message = msg
	return r
}

// Err 构造错误消息体，只有响应码，没有message
func Err() *R {
	return ErrMsg(RespCodeErr.msg)
}

// ErrMsg 构造错误消息体，可以自定义返回消息
func ErrMsg(msg string) *R {
	r := new(R)
	r.Code = RespCodeErr.code
	r.Message = msg
	return r
}

func Build(respCode RespCode) *R {
	r := new(R)
	r.Code = respCode.code
	r.Message = respCode.msg
	return r
}

func BuildData(respCode RespCode, data interface{}) *R {
	r := new(R)
	r.Code = respCode.code
	r.Message = respCode.msg
	r.Data = data
	return r
}

func ResOk(c *gin.Context) {
	r := Ok()
	c.JSON(http.StatusOK, r)
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
