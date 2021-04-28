package verifier

import "github.com/smallnest/rpcx/client"

type Handler struct {
	Con *client.XClient
}

type Args struct {
	Token  string
	Scope  string
	Issuer string
}

type Reply struct {
	Code  int8
	UUID  string
	Scope string
}
