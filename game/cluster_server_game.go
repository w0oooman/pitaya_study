package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"strings"
)

type RequestGameTest struct {
	component.Base
}

var app pitaya.Pitaya

func (m *RequestGameTest) Init()     {}
func (m *RequestGameTest) Shutdown() {}
func (m *RequestGameTest) TestEcho(ctx context.Context, in []byte) ([]byte, error) {
	fmt.Println("game RequestGameTest TestEcho...")
	return in, nil
}

func main() {
	conf := viper.New()
	conf.SetDefault("pitaya.worker.redis.url", "localhost:6379")
	conf.SetDefault("pitaya.worker.redis.pool", "3")
	conf.SetDefault("pitaya.cluster.rpc.client.nats.connect", "localhost:4222")
	conf.SetDefault("pitaya.cluster.rpc.server.nats.connect", "localhost:4222")
	conf.SetDefault("pitaya.cluster.sd.etcd.endpoints", "localhost:2379")

	config := config.NewConfig(conf)

	builder := pitaya.NewBuilderWithConfigs(false, "game", pitaya.Cluster, map[string]string{}, config)
	app = builder.Build()

	defer app.Shutdown()
	app.RegisterRemote(&RequestGameTest{},
		component.WithName("reqgametest"),
		component.WithNameFunc(strings.ToLower),
	)

	app.Start()
}
