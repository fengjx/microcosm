package api

import (
	"github.com/gin-gonic/gin"
	"microcosm/module/test"
	"net/http"
)

func TestList(c *gin.Context) {
	list, err := test.FindList()
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, list)
}
