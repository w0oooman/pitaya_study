package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/acceptorwrapper"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/logger"
	pb "pitaya_study/proto/pb/go"
	"pitaya_study/study/gate/login"
	"strings"
	"time"
)

var appGate pitaya.Pitaya

type RequestGateTest struct {
	component.Base
}

func (m *RequestGateTest) TestEcho(ctx context.Context, in *pb.TestGateRequest) (*pb.TestGateResponse, error) {
	logger.Log.Debugf("gate RequestGateTest TestEcho..., id = %d\n", in.Id)
	return &pb.TestGateResponse{Id: in.Id}, nil
}

func main() {
	conf := viper.New()
	conf.SetDefault("pitaya.worker.redis.url", "localhost:6379")
	conf.SetDefault("pitaya.worker.redis.pool", "3")
	conf.SetDefault("pitaya.cluster.rpc.client.nats.connect", "localhost:42221")
	conf.SetDefault("pitaya.cluster.rpc.server.nats.connect", "localhost:42221")
	conf.SetDefault("pitaya.cluster.sd.etcd.endpoints", "localhost:2379")
	conf.Set("pitaya.conn.ratelimiting.limit", 6)
	conf.Set("pitaya.conn.ratelimiting.interval", time.Second)

	vConfig := config.NewConfig(conf)
	acceptor := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", 55561))
	rateLimitConfig := config.NewPitayaConfig(vConfig).Conn.RateLimiting
	tcp := acceptorwrapper.WithWrappers(
		acceptor,
		acceptorwrapper.NewRateLimitingWrapper(nil, rateLimitConfig),
	)

	builder := pitaya.NewBuilderWithConfigs(true, "gate", pitaya.Cluster, map[string]string{}, vConfig)
	builder.AddAcceptor(tcp)
	appGate = builder.Build()

	appGate.Register(login.NewLogin(appGate),
		component.WithName("login"),
		component.WithNameFunc(strings.ToLower),
	)
	appGate.Register(&RequestGateTest{},
		component.WithName("reqgatetest"),
		component.WithNameFunc(strings.ToLower),
	)

	defer appGate.Shutdown()

	appGate.Start()
}
