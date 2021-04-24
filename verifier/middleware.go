package verifier

import (
	"github.com/gin-gonic/gin"
)

// Pass scope like
// 	admin; moderator; user
//	*

func Middleware(scopes string) gin.HandlerFunc {
	client := new(RPC_Handler)
	client.con = New()

	return func(c *gin.Context) {
		token := ExtractToken(c)
		if token == "" {
			AbortBadRequest(c)
			return
		}

		authorization, err := client.get_authorization(token, scopes)
		if err != nil {
			AbortServerError(c)
			return
		}
		switch authorization {
		case 0:
			c.Next()
		case 1:
			AbortUnauthorized(c)
			return
		case 2:
			AbortTokenExpired(c)
			return
		case 3:
			AbortInsufficentPermissions(c)
			return
		default:
			AbortServerError(c)
			return
		}
		c.Next()
	}
}
