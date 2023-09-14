package auth

import (
	response "05-TM-middleware-docs/pkg/web"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Token") != token {
			response.Failure(c, http.StatusUnauthorized, errors.New("token inválido"))
			return
		}

		c.Next()
	}
}
