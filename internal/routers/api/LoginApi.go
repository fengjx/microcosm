package api

import (
	"microcosm/internal/pkg/utils"
	"microcosm/internal/routers/form"
	"microcosm/internal/routers/protoc"
	"microcosm/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	loginService = service.GetLoginService()
)

// Login 后台登录
func Login(c *gin.Context) {
	var loginForm form.LoginForm
	var err error
	if err = c.ShouldBind(&loginForm); err != nil || loginForm.Username == "" {
		protoc.ResError(c, "情输入用户名和密码")
		return
	}
	token := loginService.Login(loginForm)
	if token != "" {
		data := map[string]string{"token": token}
		protoc.ResSuccess(c, data)
		return
	}
	protoc.ResError(c, "用户名密码错误")
}

func LoginInfo(c *gin.Context) {
	claims, _ := c.Get("login-info")
	jwtClaims := claims.(*utils.JwtClaims)
	data := map[string]string{"username": jwtClaims.Username}
	data["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	protoc.ResSuccess(c, data)
}

func Logout(c *gin.Context) {
	protoc.ResOk(c)
}
