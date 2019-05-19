package gate

import (
	"context"
	"fmt"
	"github.com/etcd-io/etcd/client"
	"time"
)

type gateWay struct {
	client.Client
}

func NewGateWay(endPoints []string) {
	cfg := client.Config{
		Endpoints:               endPoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	clt, err := client.New(cfg)
	if err != nil {
		fmt.Printf("new etcd client fail.err %v\n", err)
	}
	keyApi := client.NewKeysAPI(clt)
	resp, err := keyApi.Set(context.Background(), "tskey1", "ts1", nil)
	if err != nil {
		fmt.Println("set key. err:", err)
	}
	fmt.Println("resp act:", resp.Action)
}
