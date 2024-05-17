package actor

import (
	"github.com/nats-io/nats.go"
	e "github.com/topfreegames/pitaya/v2/errors"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/protos"
	"github.com/topfreegames/pitaya/v2/util"
)

type userActor struct {
	userID       int64
	ch           chan *protos.Request
	conn         *nats.Conn
	pitayaServer protos.PitayaServer
}

func newUserActor(userID int64, conn *nats.Conn, pitayaServer protos.PitayaServer, buffSize int, stopChan chan bool) *userActor {
	res := &userActor{
		userID:       userID,
		ch:           make(chan *protos.Request, buffSize),
		conn:         conn,
		pitayaServer: pitayaServer,
	}
	go res.process(stopChan)
	return res
}

func (u *userActor) ID() int64 {
	return u.userID
}

func (u *userActor) Request(req *protos.Request) {
	u.ch <- req
}

func (u *userActor) process(stopChan chan bool) {
	for {
		select {
		case req := <-u.ch:
			logger.Log.Debugf("user actor processing message %v", req.GetMsg().GetId())
			ctx, err := util.GetContextFromRequest(req)
			var response *protos.Response
			if err != nil {
				response = &protos.Response{
					Error: &protos.Error{
						Code: e.ErrInternalCode,
						Msg:  err.Error(),
					},
				}
			} else {
				response, err = u.pitayaServer.NatsCallInSingleRoutine(ctx, req)
				if err != nil {
					logger.Log.Errorf("error processing route %s: %s", req.GetMsg().GetRoute(), err)
				}
			}
			p, err := util.NatsRPCServerMarshalResponse(response)
			err = u.conn.Publish(req.GetMsg().GetReply(), p)
			if err != nil {
				logger.Log.Errorf("error sending message response: %s", err.Error())
			}
		case <-stopChan:
			return
		}
	}
}
