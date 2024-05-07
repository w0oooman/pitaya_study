package main

import (
	"fmt"
	"github.com/topfreegames/pitaya/client"
	"testing"
	"time"

	//"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/topfreegames/pitaya/conn/message"
	"github.com/topfreegames/pitaya/helpers"
)

func TestEcho(t *testing.T) {
	c := client.New(logrus.InfoLevel)

	err := c.ConnectTo(fmt.Sprintf("localhost:%d", 5555))
	if err != nil {
		fmt.Println(err)
	}
	assert.NoError(t, err)
	defer c.Disconnect()

	//uid1 := uuid.New().String()
	//_, err = c1.SendRequest("connector.testsvc.testbindid", []byte(uid1))

	//err = c.SendNotify("echo.notifytest.testecho", []byte("hello"))
	_, err = c.SendRequest("echo.reqtest.testecho", []byte("hello"))

	msg := helpers.ShouldEventuallyReceive(t, c.IncomingMsgChan, 1*time.Second).(*message.Message)
	fmt.Println("ack msg:", msg)
}
