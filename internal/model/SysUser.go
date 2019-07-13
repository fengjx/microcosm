package model

import "time"

const (
	SysUserStatusNotValid = 1
	SysUserStatusOk       = 2
	SysUserStatusDel      = 3
)

type SysUser struct {
	BaseModel
	Username     string
	Password     string
	Salt         string
	Email        string
	Mobile       string
	Status       int
	Type         int
	CreateUserId int16
	CreateTime   time.Time
}
