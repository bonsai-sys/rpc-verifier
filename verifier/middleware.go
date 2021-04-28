package verifier

import (
	"log"

	"github.com/bonsai-sys/rpc-verifier/errors"
	"github.com/gin-gonic/gin"
)

// Gin-gonic related
// 		Workflow:
//			Create Client
//			use client.get_authorization(token, scopes)
//				- remove bearer from token !
//				- scopes format "scope; scope; scope" || "*" || "scope"
//			parse reply

func Middleware(scopes string) gin.HandlerFunc {
	if Issuer == "" {
		log.Fatal("Issuer is not set. (Usage: verifier.SetIssuer(host string))")
	}
	client := new(Handler)

	return func(c *gin.Context) {
		token := ExtractToken(c)
		if token == "" {
			return
		}

		reply, err := client.Authorization(token, scopes)
		if err != nil {
			AbortUnauthorized(c, errors.ServerError())
			return
		}
		reply.Parse(c)
	}
}
