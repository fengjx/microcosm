package service

import (
	"microcosm/internal/db"
	"microcosm/internal/model"
	"microcosm/internal/pkg/utils"
	"microcosm/internal/routers/form"
	"time"
)

var sysUserService *SysUserService

func init() {
	sysUserService = new(SysUserService)
}

type SysUserService struct {
}

func (receiver *SysUserService) AddUser(data *form.AddSysUserForm) error {
	sysUser := new(model.SysUser)
	sysUser.Username = data.Username
	sysUser.Email = data.Email
	sysUser.Mobile = data.Mobile
	sysUser.CreateTime = time.Now()
	sysUser.Status = model.SysUserStatusNotValid
	sysUser.Salt = utils.RandString(4)
	sysUser.Password = utils.Md5("123456" + sysUser.Salt)
	return receiver.save(sysUser)
}

func (receiver *SysUserService) save(sysUser *model.SysUser) error {
	orm := db.GetDB()
	err := orm.Create(&sysUser).Error
	return err
}

func GetSysUserService() *SysUserService {
	return sysUserService
}
