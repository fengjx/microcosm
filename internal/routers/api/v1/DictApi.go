package v1

import (
	"microcosm/internal/routers/form"
	"microcosm/internal/routers/protoc"
	"microcosm/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	dictService = service.GetDictService()
)

// DictPageList 分页查询字典数据
func DictPageList(c *gin.Context) {
	var pageForm form.PageForm
	err := c.ShouldBind(&pageForm)
	page, err := dictService.PageList(pageForm)
	if err != nil {
		protoc.ResError(c, "查询字典数据异常")
		return
	}
	protoc.ResSuccess(c, page)
}
