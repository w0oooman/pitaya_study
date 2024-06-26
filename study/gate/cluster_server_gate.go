package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	pb "pitaya_study/proto/pb/go"
	"strings"
)

type RequestGateTest struct {
	component.Base
}

func (m *RequestGateTest) TestEcho(ctx context.Context, in *pb.TestGateRequest) (*pb.TestGateResponse, error) {
	fmt.Printf("gate RequestGateTest TestEcho..., id = %d\n", in.Id)
	return &pb.TestGateResponse{Id: in.Id}, nil
}

var appGate pitaya.Pitaya

func main() {
	conf := viper.New()
	conf.SetDefault("pitaya.worker.redis.url", "localhost:6379")
	conf.SetDefault("pitaya.worker.redis.pool", "3")
	conf.SetDefault("pitaya.cluster.rpc.client.nats.connect", "localhost:42221")
	conf.SetDefault("pitaya.cluster.rpc.server.nats.connect", "localhost:42221")
	conf.SetDefault("pitaya.cluster.sd.etcd.endpoints", "localhost:2379")

	config := config.NewConfig(conf)

	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", 55561))

	builder := pitaya.NewBuilderWithConfigs(true, "gate", pitaya.Cluster, map[string]string{}, config)
	builder.AddAcceptor(tcp)
	appGate = builder.Build()

	defer appGate.Shutdown()
	appGate.Register(&RequestGateTest{},
		component.WithName("reqgatetest"),
		component.WithNameFunc(strings.ToLower),
	)

	appGate.Start()
}
