package main

import (
	"fmt"
	"testing"
	"time"

	//"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/topfreegames/pitaya/client"
	"github.com/topfreegames/pitaya/conn/message"
	"github.com/topfreegames/pitaya/helpers"
)

func TestClusterEcho(t *testing.T) {
	c := client.New(logrus.InfoLevel)

	err := c.ConnectTo(fmt.Sprintf("localhost:%d", 5556))
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, err)
	defer c.Disconnect()

	_, err = c.SendRequest("gate.reqgatetest.testecho", []byte("hello gate"))
	msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 13*time.Second).(*message.Message)
	fmt.Println("ack msg:", msg)

	_, err = c.SendRequest("game.reqgametest.testecho", []byte("hello game"))
	msg = helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 3*time.Second).(*message.Message)
	fmt.Println("ack msg:", msg)
}
