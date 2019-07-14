package form

// PageForm 分页查询参数
type PageForm struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
