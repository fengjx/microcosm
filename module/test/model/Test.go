package model

import (
	"time"
)

type Test struct {
	Id         int32 `gorm:"column:id;primary_key:true"`
	Name       string
	Age        int
	CreateTime time.Time
}

func (t *Test) TableName() string {
	return "tb_test"
}
