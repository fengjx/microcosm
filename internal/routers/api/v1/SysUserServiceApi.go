package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"microcosm/internal/routers/data"
	"microcosm/internal/routers/form"
	"microcosm/internal/service"
)

var sysUserService = service.GetSysUserService()

func AddSysUser(c *gin.Context) {
	var formData form.AddSysUserForm
	var err error
	log.Info(c.Query("username"), c.ContentType())
	if err = c.ShouldBind(&formData); err == nil {
		err = sysUserService.AddUser(&formData)
		data.Res(c, err)
	}
}

func ListSysUser(c *gin.Context) {
}
