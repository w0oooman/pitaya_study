package main

import (
	"fmt"
	"math/rand"
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
	t.Parallel()
	c := client.New(logrus.InfoLevel)
	err := c.ConnectTo(fmt.Sprintf("localhost:%d", 55561))
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, err)
	defer c.Disconnect()

	c2 := client.New(logrus.InfoLevel)
	err = c2.ConnectTo(fmt.Sprintf("localhost:%d", 55561))
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, err)
	defer c2.Disconnect()

	callGame := func(c *client.Client) {
		gameRequest := &pb.TestGameRequest{Id: 456}
		gameRequestByte, err := proto.Marshal(gameRequest)
		assert.NoError(t, err)
		_, err = c.SendRequest("game.reqgametest.testecho", gameRequestByte)
		assert.NoError(t, err)
		msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 3*time.Second).(*message.Message)
		gameResponse := &pb.TestGateResponse{}
		err = proto.Unmarshal(msg.Data, gameResponse)
		assert.NoError(t, err)
		fmt.Printf("game response:%+v\n", gameResponse)
	}

	callGuild := func(c *client.Client) {
		gameRequest := &pb.TestGameRequest{Id: 888}
		gameRequestByte, err := proto.Marshal(gameRequest)
		assert.NoError(t, err)
		_, err = c.SendRequest("guild.reqgametest.testecho", gameRequestByte)
		assert.NoError(t, err)
		msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 3*time.Second).(*message.Message)
		gameResponse := &pb.TestGateResponse{}
		err = proto.Unmarshal(msg.Data, gameResponse)
		assert.NoError(t, err)
		fmt.Printf("guild response:%+v\n", gameResponse)
	}

	callLogin := func(c *client.Client) {
		gameRequest := &pb.LoginRequest{RoleId: 578 + rand.Int63n(1000)}
		gameRequestByte, err := proto.Marshal(gameRequest)
		assert.NoError(t, err)
		_, err = c.SendRequest("gate.login.login", gameRequestByte)
		assert.NoError(t, err)
		msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 3*time.Second).(*message.Message)
		gameResponse := &pb.LoginResponse{}
		err = proto.Unmarshal(msg.Data, gameResponse)
		assert.NoError(t, err)
		fmt.Printf("login response:%+v\n", gameResponse)
	}

	// failed
	//callGame(c)
	//callGuild(c)

	// login
	callLogin(c)
	callLogin(c2)

	time.Sleep(500 * time.Millisecond)
	callGame(c)
	callGame(c2)
	time.Sleep(500 * time.Millisecond)
	callGuild(c)
	callGuild(c2)
	time.Sleep(500 * time.Millisecond)

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

	time.Sleep(5 * time.Second)
}
