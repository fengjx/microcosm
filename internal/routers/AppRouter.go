package routers

import (
	"microcosm/internal/config"
	"microcosm/internal/middleware/jwt"
	"microcosm/internal/routers/api"
	v1 "microcosm/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AppRouter struct {
}

func (r *AppRouter) Start(config *config.Config, engine *gin.Engine) {
	apiv1 := engine.Group("/api/v1")
	apiv1.Use(jwt.JwtHandler())
	{
		apiv1.POST("/sys/user/add", v1.AddSysUser)
		apiv1.GET("/sys/user/list", v1.ListSysUser)
	}
	engine.POST("/login", api.Login)
	engine.GET("/logout", api.Logout).Use(jwt.JwtHandler())
	engine.GET("/login-info", api.LoginInfo).Use(jwt.JwtHandler())
	engine.NoRoute(api.Index)
	engine.GET("/hello", api.Hello)
	engine.GET("/test/list", api.TestList)
	log.Info("Approuter started")
}

func (r *AppRouter) Stop(config *config.Config, engine *gin.Engine) {
	log.Info("Approuter stop")
}

func NewAppRouter() *AppRouter {
	return new(AppRouter)
}
