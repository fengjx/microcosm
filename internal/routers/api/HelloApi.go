package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello microcosm \n")
}
