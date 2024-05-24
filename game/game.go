package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/logger"
	pb "pitaya_study/proto/pb/go"
	"strings"
)

type RequestGameTest struct {
	component.Base
}

var app pitaya.Pitaya

func (m *RequestGameTest) TestEcho(ctx context.Context, in *pb.TestGameRequest) (*pb.TestGameResponse, error) {
	logger.Log.Debugf("game RequestGameTest TestEcho..., id = %d\n", in.Id)
	return &pb.TestGameResponse{Id: in.Id}, nil
}

func main() {
	conf := viper.New()
	conf.SetDefault("pitaya.worker.redis.url", "localhost:6379")
	conf.SetDefault("pitaya.worker.redis.pool", "3")
	conf.SetDefault("pitaya.cluster.rpc.client.nats.connect", "localhost:42221")
	conf.SetDefault("pitaya.cluster.rpc.server.nats.connect", "localhost:42221")
	conf.SetDefault("pitaya.cluster.sd.etcd.endpoints", "localhost:2379")

	config := config.NewConfig(conf)

	builder := pitaya.NewBuilderWithConfigs(false, "game", pitaya.Cluster, map[string]string{}, config)
	app = builder.Build()
	builder.RPCServer.SetRPCServiceUserActor()

	defer app.Shutdown()
	app.Register(&RequestGameTest{},
		component.WithName("reqgametest"),
		component.WithNameFunc(strings.ToLower),
	)

	app.Start()
}
