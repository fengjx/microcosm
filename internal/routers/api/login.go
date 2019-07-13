package api

import (
	"github.com/gin-gonic/gin"
	"microcosm/internal/routers/form"
	"microcosm/internal/routers/protoc"
	"microcosm/internal/service"
)

var (
	loginService = service.GetLoginService()
)

func Login(c *gin.Context) {
	var loginForm form.LoginForm
	var err error
	if err = c.ShouldBind(&loginForm); err != nil {
		protoc.ErrMsg("情输入用户名和密码")
	}
	token := loginService.Login(loginForm)
	if token != "" {
		data := map[string]string{"token": token}
		protoc.ResSuccess(c, data)
		return
	}
	protoc.ResError(c, "用户名密码错误")
}
