package routers

import (
	"microcosm/internal/config"
	"microcosm/internal/middleware/jwt"
	"microcosm/internal/routers/api"
	v1 "microcosm/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// AppRouter 路由定义
type AppRouter struct {
}

// Start 启动回调函数
func (r *AppRouter) Start(config *config.Config, engine *gin.Engine) {
	apiv1 := engine.Group("/api/v1")
	apiv1.Use(jwt.JwtHandler())
	{
		apiv1.POST("/sys/user/add", v1.AddUser)
		apiv1.GET("/sys/user/list", v1.ListUser)
		apiv1.GET("/sys/dict/page-list", v1.DictPageList)
	}
	authAPI := engine.Group("/")
	authAPI.Use(jwt.JwtHandler())
	{
		authAPI.GET("/logout", api.Logout).Use(jwt.JwtHandler())
		authAPI.GET("/login-info", api.LoginInfo).Use(jwt.JwtHandler())
	}
	engine.GET("/config", api.Config)
	engine.POST("/login", api.Login)
	engine.GET("/hello", api.Hello)
	log.Info("Approuter started")
}

// Stop AppRouter注销回调函数
func (r *AppRouter) Stop(config *config.Config, engine *gin.Engine) {
	log.Info("Approuter stop")
}

// NewAppRouter 创建一个AppRouter实例
func NewAppRouter() *AppRouter {
	return new(AppRouter)
}
