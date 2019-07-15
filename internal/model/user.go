package model

const (
	UserStatusNotValid = 1
	UserStatusOk       = 2
	UserStatusDel      = 3
)

type User struct {
	BaseModel
	CreateUserID int64  `gorm:"column:create_user_id"`
	Email        string `gorm:"column:email"`
	Mobile       string `gorm:"column:mobile"`
	Password     string `gorm:"column:password"`
	Salt         string `gorm:"column:salt"`
	Status       int    `gorm:"column:status"`
	Type         int    `gorm:"column:type"`
	Username     string `gorm:"column:username"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
