package main

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/serialize/json"
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

func main() {
	defer pitaya.Shutdown()
	pitaya.SetSerializer(json.NewSerializer())

	pitaya.Register(&RequestTest{},
		component.WithName("reqtest"),
		component.WithNameFunc(strings.ToLower),
	)

	pitaya.Configure(true, "echo", pitaya.Standalone, map[string]string{})
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf("%s:%d", "localhost", 5555))
	pitaya.AddAcceptor(tcp)
	pitaya.Start()
}
