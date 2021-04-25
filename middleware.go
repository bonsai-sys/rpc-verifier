package verifier

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Pass scope like
// 	admin; moderator; user
//	*

func Middleware(scopes string) gin.HandlerFunc {
	if Issuer == "" {
		log.Fatal("Issuer not set. (Usage verifier.SetIssuer())")
	}
	client := new(RPC_Handler)
	client.con = New()

	return func(c *gin.Context) {
		token := ExtractToken(c)
		if token == "" {
			AbortBadRequest(c)
			return
		}

		reply, err := client.get_authorization(token, scopes)
		if err != nil {
			AbortServerError(c)
			return
		}

		switch reply.Code {
		case 0:
			c.Set("UUID", reply.UUID)
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
		case 4:
			AbortWrongIssuer(c)
			return
		default:
			AbortServerError(c)
			return
		}
	}
}
