package service

import (
	"microcosm/internal/db"
	"microcosm/internal/model"
	"microcosm/internal/pkg/utils"
	"microcosm/internal/routers/form"
	"microcosm/internal/routers/protoc"
)

var dictServiceSingleton *DictService

func init() {
	dictServiceSingleton = new(DictService)
}

// GetDictService 获得DictService实例
func GetDictService() *DictService {
	return dictServiceSingleton
}

// DictService 字典管理
type DictService struct {
}

// PageList 分页查询字典信息
func (dictService *DictService) PageList(pageForm form.PageForm) (*protoc.Page, error) {
	var total int
	var dicts []*model.Dict
	offset, limit := utils.GetOffsetLimit(pageForm.Page, pageForm.Size)
	errors := db.GetOrm().Model(&model.Dict{}).Count(&total).Offset(offset).Limit(limit).Find(&dicts).GetErrors()
	if len(errors) > 0 {
		return nil, errors[0]
	}
	dictPage := protoc.BuildPage(total, pageForm.Size, pageForm.Page, dicts)
	return dictPage, nil
}
