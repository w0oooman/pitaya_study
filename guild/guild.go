package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/logger"
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

func newGuildTest(app pitaya.Pitaya) *RequestGuildTest {
	return &RequestGuildTest{
		app: app,
	}
}

type RequestGuildTest struct {
	component.Base
	app pitaya.Pitaya
}

var app pitaya.Pitaya

func (m *RequestGuildTest) TestEcho(ctx context.Context, in *pb.TestGameRequest) (*pb.TestGameResponse, error) {
	logger.Log.Debugf("guild RequestGameTest TestEcho..., id = %d, routineID=%d\n", in.Id, getRoutineID())
	s := m.app.GetSessionFromCtx(ctx)
	if s == nil {
		return nil, fmt.Errorf("session is nil")
	}
	uid := s.UID()
	if uid == "" {
		return nil, fmt.Errorf("session uid is empty")
	}
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
	app = builder.Build()

	defer app.Shutdown()
	app.Register(newGuildTest(app),
		component.WithName("reqgametest"),
		component.WithNameFunc(strings.ToLower),
	)

	app.Start()
}
