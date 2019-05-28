package api

import (
	"github.com/gin-gonic/gin"
	"microcosm/internal/service"
	"net/http"
)

var testService = service.GetTestService()

func TestList(c *gin.Context) {
	list, err := testService.FindList()
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, list)
}
