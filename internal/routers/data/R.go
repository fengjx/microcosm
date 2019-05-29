package data

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const codeOk = 1
const codeErr = 0
const msgOk = "success"
const msgErr = "error"

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok() *R {
	return OkMsg(msgOk)
}

func OkMsg(msg string) *R {
	r := new(R)
	r.Code = codeOk
	r.Msg = msg
	return r
}

func Err() *R {
	return OkMsg(msgErr)
}

func ErrMsg(msg string) *R {
	r := new(R)
	r.Code = codeErr
	r.Msg = msg
	return r
}

func Res(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Ok())
	} else {
		c.JSON(http.StatusInternalServerError, Ok())
	}
}
