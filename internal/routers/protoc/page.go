package protoc

import "microcosm/internal/pkg/utils"

// Page 是一个分页对象，封装分页数据返回给前端
type Page struct {
	Total       int           `json:"total"`
	Size        int           `json:"size"`
	CurPage     int           `json:"curPage"`
	TotalPage   int           `json:"totalPage"`
	HasNextPage bool          `json:"hasNextPage"`
	DataList    []interface{} `json:"dataList"`
}

// BuildPage 方法创建 一个分页对象
func BuildPage(total int, size int, curPage int, list interface{}) (page *Page) {
	dataList, ok := utils.ConvToSlice(list)
	page = new(Page)
	if !ok {
		return
	}
	page.Total = total
	page.Size = size
	page.DataList = dataList
	page.CurPage = curPage
	page.TotalPage = getTotalPage(total, size)
	page.HasNextPage = hashNest(total, size, curPage)
	return page
}

func getTotalPage(total int, size int) int {
	return (total - 1) / size
}

func hashNest(total int, size int, curPage int) bool {
	return getTotalPage(total, size) > curPage
}
