package model

// BaseModel 实体类父类
type BaseModel struct {
	ID         int64                    `gorm:"column:id;primary_key:true" json:"id"`
	CreateTime TimeImplementedMarshaler `gorm:"column:create_time" json:"createTime"`
}
