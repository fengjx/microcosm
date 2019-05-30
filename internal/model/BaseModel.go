package model

type BaseModel struct {
	Id uint32 `gorm:"column:id;primary_key:true"`
}
