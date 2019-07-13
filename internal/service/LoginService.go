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

func (receiver *LoginService) Login(form form.LoginForm) (token string) {
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
