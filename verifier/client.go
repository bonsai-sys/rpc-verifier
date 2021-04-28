package verifier

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/c3b5aw/go-utils/env"
	"github.com/smallnest/rpcx/client"
)

var Issuer string

func SetIssuer(iss string) {
	Issuer = iss
}

func Opts() client.Option {
	opts := client.DefaultOption
	opts.Heartbeat = true
	opts.HeartbeatInterval = 3 * time.Second
	opts.MaxWaitForHeartbeat = 5 * time.Second
	opts.IdleTimeout = 15 * time.Second
	if env.GBool("RPC_VERIFIER_SSL", false) {
		opts.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return opts
}

func New() *Handler {
	_handler := new(Handler)
	disc, err := client.NewPeer2PeerDiscovery("tcp@"+env.GString("RPC_VERIFIER_ADDR", "verify.bonsai-sys.io:8792"), "")
	if err != nil {
		log.Fatal("Unable to reach verify server")
	}
	xClient := client.NewXClient("auth", client.Failtry, client.RandomSelect, disc, Opts())
	_handler.Con = &xClient
	return _handler
}
