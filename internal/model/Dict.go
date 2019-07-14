package model

type Dict struct {
	BaseModel
	GroupLabel string `gorm:"column:group_label"`
	GroupVal   string `gorm:"column:group_val"`
	Label      string `gorm:"column:label"`
	Val        int    `gorm:"column:val"`
}

// TableName sets the insert table name for this struct type
func (d *Dict) TableName() string {
	return "dict"
}
