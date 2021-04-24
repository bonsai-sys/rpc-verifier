package verifier

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/c3b5aw/go-utils/env"
	"github.com/smallnest/rpcx/client"
)

var AuthClient client.XClient = New()

func New() client.XClient {
	disc, err := client.NewPeer2PeerDiscovery("tcp@"+env.String("BONSAI_VERIFY_ADDR", "localhost:8972"), "")
	if err != nil {
		log.Fatal("Unable to reach verify server")
	}
	options := client.DefaultOption

	options.Heartbeat = true
	options.HeartbeatInterval = 3 * time.Second
	options.MaxWaitForHeartbeat = 5 * time.Second
	options.IdleTimeout = 15 * time.Second
	if env.Bool("BONSAI_VERIFIER_SSL", false) {
		conf := &tls.Config{
			InsecureSkipVerify: true,
		}
		options.TLSConfig = conf
	}

	xClient := client.NewXClient(
		"auth",
		client.Failtry,
		client.RandomSelect,
		disc,
		options,
	)
	return xClient
}

func Verify(token, scopes string) (bool, error) {
	reply := new(Reply)
	args := Args{
		Token: token,
		Scope: scopes,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()

	err := AuthClient.Call(ctx, "Verify", args, &reply)
	if err != nil {
		return false, CommunicationError()
	}

	if !reply.Valid {
		return false, nil
	}
	return true, nil
}
