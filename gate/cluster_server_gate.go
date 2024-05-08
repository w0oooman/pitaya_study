package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"strings"
)

type RequestGateTest struct {
	component.Base
}

func (m *RequestGateTest) Init()     {}
func (m *RequestGateTest) Shutdown() {}
func (m *RequestGateTest) TestEcho(ctx context.Context, in []byte) ([]byte, error) {
	fmt.Println("gate RequestGateTest TestEcho...")
	return in, nil
}

var app pitaya.Pitaya

func main() {
	conf := viper.New()
	conf.SetDefault("pitaya.worker.redis.url", "localhost:6379")
	conf.SetDefault("pitaya.worker.redis.pool", "3")
	conf.SetDefault("pitaya.cluster.rpc.client.nats.connect", "localhost:4222")
	conf.SetDefault("pitaya.cluster.rpc.server.nats.connect", "localhost:4222")
	conf.SetDefault("pitaya.cluster.sd.etcd.endpoints", "localhost:2379")

	config := config.NewConfig(conf)

	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", 55561))

	builder := pitaya.NewBuilderWithConfigs(true, "gate", pitaya.Cluster, map[string]string{}, config)
	builder.AddAcceptor(tcp)
	app = builder.Build()

	defer app.Shutdown()
	app.RegisterRemote(&RequestGateTest{},
		component.WithName("RequestGateTest"),
		component.WithNameFunc(strings.ToLower),
	)

	app.Start()
}
