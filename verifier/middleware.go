package verifier

import (
	"github.com/gin-gonic/gin"
)

// Pass scope like
// 	admin; moderator; user

func Middleware(scopes string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ExtractToken(c)
		if token == "" {
			AbortBadRequest(c)
			return
		}

		authorized, err := Verify(token, scopes)
		if err != nil {
			AbortServerError(c)
			return
		}
		if !authorized {
			AbortUnauthorized(c)
			return
		}
		c.Next()
	}
}
