package verifier

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortUnauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":  http.StatusUnauthorized,
		"error": "unauthorized",
	})
}

func AbortServerError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code":  http.StatusInternalServerError,
		"error": "server error",
	})
}

func AbortBadRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code":  http.StatusBadRequest,
		"error": "bad request",
	})
}

func AbortInsufficentPermissions(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":  http.StatusUnauthorized,
		"error": "insufficent permissions",
	})
}

func AbortTokenExpired(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":  http.StatusUnauthorized,
		"error": "token expired",
	})
}

func ExtractToken(c *gin.Context) string {
	t := c.Request.Header.Get("Authorization")
	if len(t) < 128 {
		return ""
	}
	return t[7:]
}
