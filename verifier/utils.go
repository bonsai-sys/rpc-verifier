package verifier

import (
	"net/http"

	"github.com/bonsai-sys/rpc-verifier/errors"
	"github.com/c3b5aw/go-utils/env"
	"github.com/gin-gonic/gin"
)

var DefaultError bool = env.GBool("RPC_VERIFIER_ERROR_DEFAULT", false)

func AbortUnauthorized(c *gin.Context, m string) {
	if DefaultError {
		m = errors.Unauthorized()
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":  http.StatusUnauthorized,
		"error": m,
	})
}

func ExtractToken(c *gin.Context) string {
	t := c.Request.Header.Get("Authorization")
	if len(t) < 128 {
		AbortUnauthorized(c, errors.BadFormat())
		return ""
	}
	return t[7:]
}
