package actor

import (
	"github.com/topfreegames/pitaya/v2/protos"
)

type Actors interface {
	Request(int64, *protos.Request)
}
