package app

import (
	"github.com/gin-gonic/gin"
	"microcosm/conf"
)

type RouteStarter interface {
	Start(config *conf.Config, engine *gin.Engine)
	Stop(config *conf.Config, engine *gin.Engine)
}
