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

type NotifyTest struct {
	component.Base
}

func (m *NotifyTest) Init()     {}
func (m *NotifyTest) Shutdown() {}
func (m *NotifyTest) TestEcho(ctx context.Context) {
	fmt.Println("NotifyTest TestEcho...")
}

type RequestTest struct {
	component.Base
}

func (m *RequestTest) Init()     {}
func (m *RequestTest) Shutdown() {}
func (m *RequestTest) TestEcho(ctx context.Context) ([]byte, error) {
	fmt.Println("RequestTest TestEcho...")
	return []byte{0x30, 0x31}, nil
}

var appEcho pitaya.Pitaya

func main() {
	conf := viper.New()
	conf.SetDefault("pitaya.worker.redis.url", "localhost:6379")
	conf.SetDefault("pitaya.worker.redis.pool", "3")
	conf.SetDefault("pitaya.cluster.rpc.client.nats.connect", "localhost:4222")
	conf.SetDefault("pitaya.cluster.rpc.server.nats.connect", "localhost:4222")
	conf.SetDefault("pitaya.cluster.sd.etcd.endpoints", "localhost:2379")

	config := config.NewConfig(conf)

	builder := pitaya.NewBuilderWithConfigs(false, "echo", pitaya.Cluster, map[string]string{}, config)
	appEcho = builder.Build()
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf("%s:%d", "localhost", 5555))
	builder.AddAcceptor(tcp)

	defer appEcho.Shutdown()
	appEcho.RegisterRemote(&RequestTest{},
		component.WithName("reqtest"),
		component.WithNameFunc(strings.ToLower),
	)

	appEcho.Start()
}
