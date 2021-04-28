package verifier

import (
	"context"
	"time"

	"github.com/bonsai-sys/rpc-verifier/errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Authorization(token, scopes string) (*Reply, error) {
	resp := new(Reply)
	args := Args{
		Token:  token,
		Scope:  scopes,
		Issuer: Issuer,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()

	err := (*h.Con).Call(ctx, "Verify", args, &resp)
	if err != nil {
		return nil, errors.CommunicationError()
	}
	return resp, nil
}

func (r *Reply) Parse(c *gin.Context) {
	switch r.Code {
	case 0:
		c.Set("UUID", r.UUID)
		c.Next()
	case 1:
		AbortUnauthorized(c, errors.Unauthorized())
	case 2:
		AbortUnauthorized(c, errors.TokenExpired())
	case 3:
		AbortUnauthorized(c, errors.InsufficentPermissions())
	case 4:
		AbortUnauthorized(c, errors.WrongIssuer())
	default:
		AbortUnauthorized(c, errors.ServerError())
	}
}
