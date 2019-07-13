package service

import (
	"github.com/jinzhu/gorm"
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
	orm *gorm.DB
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

func (receiver SysUserService) UpdateSysUser(sysUser *model.SysUser) error {
	return receiver.orm.Save(sysUser).Error
}

func (receiver *SysUserService) save(sysUser *model.SysUser) error {
	orm := db.GetDB()
	err := orm.Create(&sysUser).Error
	return err
}

func (receiver SysUserService) DeleteSysUser(id uint32) error {
	user := new(model.SysUser)
	user.Id = id
	return receiver.orm.Delete(user).Error
}

func (receiver *SysUserService) GetSysUserById(id uint32) model.SysUser {
	var user model.SysUser
	receiver.orm.First(&user, id)
	return user
}

func (receiver SysUserService) FindSysUserList(offset int8, limit uint, orderBy string) ([]model.SysUser,  error) {
	var users []model.SysUser
	err := receiver.orm.Offset(offset).Limit(limit).Order(orderBy).Find(&users).Error
	return users, err
}

func GetSysUserService() *SysUserService {
	return sysUserService
}
