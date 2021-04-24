package verifier

import "github.com/smallnest/rpcx/client"

type RPC_Handler struct {
	con *client.XClient
}

type Args struct {
	Token string
	Scope string
}

type Reply struct {
	Code int8
}
