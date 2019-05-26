package routers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"microcosm/conf"
	"microcosm/routers/api"
)

type AppRouter struct {
}

func (r *AppRouter) Start(config *conf.Config, engine *gin.Engine) {
	engine.GET("/hello", api.Hello)
	engine.GET("/test/list", api.TestList)
	log.Info("Approuter started")
}

func (r *AppRouter) Stop(config *conf.Config, engine *gin.Engine) {

	log.Info("Approuter stop")
}

func NewAppRouter() *AppRouter {
	return new(AppRouter)
}
