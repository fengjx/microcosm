package app

import (
	"github.com/gin-gonic/gin"
	"microcosm/internal/config"
)

type RouteStarter interface {
	Start(config *config.Config, engine *gin.Engine)
	Stop(config *config.Config, engine *gin.Engine)
}
