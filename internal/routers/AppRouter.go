package routers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"microcosm/internal/config"
	"microcosm/internal/routers/api"
	"microcosm/internal/routers/api/v1"
)

type AppRouter struct {
}

func (r *AppRouter) Start(config *config.Config, engine *gin.Engine) {
	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"fjx": "123",
	}))
	{
		authorized.POST("/sys/user/add", v1.AddSysUser)
		authorized.GET("/sys/user/list", v1.ListSysUser)
	}

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
