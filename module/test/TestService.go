package test

import (
	"github.com/jinzhu/gorm"
	"microcosm/db"
	"microcosm/module/test/model"
)

func FindList() ([]*model.Test, error) {
	var list []*model.Test
	err := db.GetDB().Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}
