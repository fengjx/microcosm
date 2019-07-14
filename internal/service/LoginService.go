package service

import (
	"microcosm/internal/config"
	"microcosm/internal/pkg/utils"
	"microcosm/internal/routers/form"
)

var loginService *LoginService

func init() {
	loginService = new(LoginService)
}

type LoginService struct {
}

// Login 校验用户名密码
func (receiver *LoginService) Login(form form.LoginForm) (token string) {
	if form.Username == "" {
		return ""
	}
	users := config.GetConfig().GetUsers()
	pwd := users[form.Username]
	if pwd == form.Password {
		token, err := utils.GenJwtToken(form.Username, form.Password)
		if err == nil {
			return token
		}
	}
	return ""
}

func GetLoginService() *LoginService {
	return loginService
}
