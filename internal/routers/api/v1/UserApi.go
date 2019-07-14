package v1

import (
	"microcosm/internal/routers/form"
	"microcosm/internal/routers/protoc"
	"microcosm/internal/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	userService = service.GetUserService()
)

func AddUser(c *gin.Context) {
	var formData form.AddUserForm
	var err error
	log.Info(c.Query("username"), c.ContentType())
	if err = c.ShouldBind(&formData); err == nil {
		err = userService.AddUser(&formData)
		protoc.Res(c, err)
	}
}

func ListUser(c *gin.Context) {
}
