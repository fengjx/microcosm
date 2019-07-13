package jwt

import (
	"microcosm/internal/pkg/utils"
	"microcosm/internal/routers/protoc"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// JWT is jwt middleware
func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-token")
		if token == "" {
			protoc.ResCode(c, protoc.RespCodeNotAuth, http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims, err := utils.ParseJwtToken(token)
		if err != nil {
			log.Warn("jwt parse error", err)
			protoc.ResCode(c, protoc.RespCodeNotAuth, http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("login-info", claims)
		c.Next()
	}
}
