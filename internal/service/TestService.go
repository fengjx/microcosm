package service

import (
	"github.com/jinzhu/gorm"
	"microcosm/internal/db"
	"microcosm/internal/model"
)

var testService *TestService

func init() {
	testService = new(TestService)
}

type TestService struct {
}

func (service *TestService) FindList() ([]*model.Test, error) {
	var list []*model.Test
	err := db.GetDB().Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}

func GetTestService() *TestService {
	return testService
}
