package service

import (
	"microcosm/internal/db"
	"microcosm/internal/model"
	"microcosm/internal/pkg/utils"
	"microcosm/internal/routers/form"
	"time"
)

var userServiceSingleton *UserService

func init() {
	userServiceSingleton = new(UserService)
}

func GetUserService() *UserService {
	return userServiceSingleton
}

type UserService struct {
}

func (userService *UserService) AddUser(data *form.AddUserForm) error {
	user := new(model.User)
	user.Username = data.Username
	user.Email = data.Email
	user.Mobile = data.Mobile
	user.CreateTime = time.Now()
	user.Status = model.UserStatusNotValid
	user.Salt = utils.RandString(4)
	user.Password = utils.Md5("123456" + user.Salt)
	return userService.save(user)
}

func (userService *UserService) UpdateUser(user *model.User) error {
	return db.GetOrm().Save(user).Error
}

func (userService *UserService) save(user *model.User) error {
	orm := db.GetOrm()
	err := orm.Create(&user).Error
	return err
}

// DeleteUser 根据id删除用户
func (userService *UserService) DeleteUser(id int64) error {
	user := new(model.User)
	user.Id = id
	return db.GetOrm().Delete(user).Error
}

// GetUserById 通过id查询用户
func (userService *UserService) GetUserById(id int64) *model.User {
	var user model.User
	db.GetOrm().First(&user, id)
	return &user
}

func (userService *UserService) FindUserList(offset int8, limit uint, orderBy string) ([]model.User, error) {
	var users []model.User
	err := db.GetOrm().Offset(offset).Limit(limit).Order(orderBy).Find(&users).Error
	return users, err
}
