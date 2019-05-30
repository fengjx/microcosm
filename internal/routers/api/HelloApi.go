package api

import (
	"github.com/gin-gonic/gin"
	"microcosm/internal/config"
	"net/http"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello microcosm \n")
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", config.GetConfig().GetAppConfig())
}
