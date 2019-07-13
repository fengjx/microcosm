package jwt

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"microcosm/internal/pkg/utils"
	"microcosm/internal/routers/protoc"
	"net/http"
)

// JWT is jwt middleware
func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err == nil || token == "" {
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
		log.Info("user login: %s", claims.Username)
		c.Next()
	}
}
