package actor

import (
	"github.com/topfreegames/pitaya/v2/protos"
)

type Actor interface {
	ID() int64
	Request(*protos.Request)
}
