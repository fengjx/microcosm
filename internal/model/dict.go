package model

type Dict struct {
	BaseModel
	GroupLabel string `gorm:"column:group_label" json:"groupLabel"`
	GroupVal   string `gorm:"column:group_val" json:"groupVal"`
	Label      string `gorm:"column:label" json:"label"`
	Val        int    `gorm:"column:val" json:"val"`
}

// TableName sets the insert table name for this struct type
func (d *Dict) TableName() string {
	return "dict"
}
