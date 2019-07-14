package model

import "time"

type BaseModel struct {
	Id         int64     `gorm:"column:id;primary_key:true" json:"-"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`
}
