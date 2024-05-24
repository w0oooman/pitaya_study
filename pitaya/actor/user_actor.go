package actor

import (
	"github.com/nats-io/nats.go"
	e "github.com/topfreegames/pitaya/v2/errors"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/protos"
	"github.com/topfreegames/pitaya/v2/util"
)

const UserActorBuffSize = 16

type userActor struct {
	userID       int64
	ch           chan *protos.Request
	pitayaServer protos.PitayaServer
	conn         *nats.Conn
	stopChan     chan bool
}

func newUserActor(userID int64, conn *nats.Conn, pitayaServer protos.PitayaServer, buffSize int, stopChan chan bool) *userActor {
	logger.Log.Debugf("new user actor.userID=%d", userID)
	res := &userActor{
		userID:       userID,
		ch:           make(chan *protos.Request, buffSize),
		pitayaServer: pitayaServer,
		conn:         conn,
		stopChan:     stopChan,
	}
	go res.process()
	return res
}

func (u *userActor) ID() int64 {
	return u.userID
}

func (u *userActor) Request(req *protos.Request) {
	select {
	case u.ch <- req:
	default:
		logger.Log.Warnf("user actor channel is full, waiting push message")
		go func() {
			select {
			case u.ch <- req:
			case <-u.stopChan:
			}
		}()
	}
}

func (u *userActor) process() {
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
		case <-u.stopChan:
			return
		}
	}
}

func (u *userActor) Close() {
	// TODO:
}
