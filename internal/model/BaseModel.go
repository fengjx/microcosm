package model

type BaseModel struct {
	Id int32 `gorm:"column:id;primary_key:true"`
}
