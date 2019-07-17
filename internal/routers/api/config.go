package api

import (
	"microcosm/internal/dict"
	"microcosm/internal/routers/protoc"

	"github.com/gin-gonic/gin"
)

// Config 获得服务器配置
func Config(c *gin.Context) {
	dicts := dict.GetAllDicts()
	data := map[string]interface{}{}
	data["dicts"] = dicts
	protoc.ResSuccess(c, data)
}
