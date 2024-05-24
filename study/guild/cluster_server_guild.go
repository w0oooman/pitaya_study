package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	pb "pitaya_study/proto/pb/go"
	"runtime"
	"strconv"
	"strings"
)

func getRoutineID() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	goidStr := strings.TrimPrefix(string(b), "goroutine ")
	goidStr = goidStr[:strings.Index(goidStr, " ")]
	gid, err := strconv.ParseInt(goidStr, 10, 64)
	if err != nil {
		return -1
	}
	return gid
}

type RequestGuildTest struct {
	component.Base
}

var appGame pitaya.Pitaya

func (m *RequestGuildTest) TestEcho(ctx context.Context, in *pb.TestGameRequest) (*pb.TestGameResponse, error) {
	fmt.Printf("guild RequestGameTest TestEcho..., id = %d, routineID=%d\n", in.Id, getRoutineID())
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

	builder := pitaya.NewBuilderWithConfigs(false, "guild", pitaya.Cluster, map[string]string{}, config)
	builder.RPCServer.SetRPCServiceSingleRoutine()
	appGame = builder.Build()

	defer appGame.Shutdown()
	appGame.Register(&RequestGuildTest{},
		component.WithName("reqgametest"),
		component.WithNameFunc(strings.ToLower),
	)

	appGame.Start()
}
