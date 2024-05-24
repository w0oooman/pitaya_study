package actor

import (
	"github.com/nats-io/nats.go"
	"github.com/topfreegames/pitaya/v2/protos"
)

type userActors struct {
	actors       map[int64]*userActor
	pitayaServer protos.PitayaServer
	conn         **nats.Conn
	buffSize     int
	stopChan     chan bool
}

func NewUserActors(conn **nats.Conn, pitayaServer protos.PitayaServer, buffSize int, stopChan chan bool) *userActors {
	return &userActors{
		actors:       make(map[int64]*userActor),
		pitayaServer: pitayaServer,
		conn:         conn,
		buffSize:     buffSize,
		stopChan:     stopChan,
	}
}

func (u *userActors) Request(id int64, req *protos.Request) {
	a, ok := u.actors[id]
	if !ok {
		a = newUserActor(id, *u.conn, u.pitayaServer, u.buffSize, u.stopChan)
		u.actors[id] = a
	}
	a.Request(req)
}

func (u *userActors) Close(id int64) {
	actor, ok := u.actors[id]
	if !ok {
		return
	}
	delete(u.actors, id)
	actor.Close()
}
