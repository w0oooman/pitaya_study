package main

import (
	"fmt"
	pb "pitaya_study/proto/pb/go"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	//"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/topfreegames/pitaya/v2/client"
	"github.com/topfreegames/pitaya/v2/conn/message"
	"github.com/topfreegames/pitaya/v2/helpers"
)

func TestClusterEcho(t *testing.T) {
	c := client.New(logrus.InfoLevel)

	err := c.ConnectTo(fmt.Sprintf("localhost:%d", 55561))
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, err)
	defer c.Disconnect()

	gateRequest := &pb.TestGateRequest{Id: 123}
	gateRequestByte, err := proto.Marshal(gateRequest)
	assert.NoError(t, err)
	_, err = c.SendRequest("gate.reqgatetest.testecho", gateRequestByte)
	msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 13*time.Second).(*message.Message)
	fmt.Println("ack msg:", msg)

	_, err = c.SendRequest("game.reqgametest.testecho", []byte("hello game"))
	msg = helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 3*time.Second).(*message.Message)
	fmt.Println("ack msg:", msg)
}
