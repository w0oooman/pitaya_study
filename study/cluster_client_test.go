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

func TestClusterGateGame(t *testing.T) {
	c := client.New(logrus.InfoLevel)

	err := c.ConnectTo(fmt.Sprintf("localhost:%d", 55561))
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, err)
	defer c.Disconnect()

	go func() {
		for i := 0; i < 10; i++ {
			gateRequest := &pb.TestGateRequest{Id: 123}
			gateRequestByte, err := proto.Marshal(gateRequest)
			assert.NoError(t, err)
			_, err = c.SendRequest("gate.reqgatetest.testecho", gateRequestByte)
			msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 13*time.Second).(*message.Message)
			gateResponse := &pb.TestGateResponse{}
			err = proto.Unmarshal(msg.Data, gateResponse)
			assert.NoError(t, err)
			fmt.Printf("gate response:%+v\n", gateResponse)
		}
	}()

	gameRequest := &pb.TestGateRequest{Id: 456}
	gameRequestByte, err := proto.Marshal(gameRequest)
	assert.NoError(t, err)
	_, err = c.SendRequest("game.reqgametest.testecho", gameRequestByte)
	msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 3*time.Second).(*message.Message)
	gameResponse := &pb.TestGateResponse{}
	err = proto.Unmarshal(msg.Data, gameResponse)
	assert.NoError(t, err)
	fmt.Printf("game response:%+v\n", gameResponse)

	time.Sleep(2 * time.Second)
}
