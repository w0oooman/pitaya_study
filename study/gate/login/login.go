package login

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	pb "pitaya_study/proto/pb/go"
	"strconv"
)

type Login struct {
	component.Base
	app pitaya.Pitaya
}

func NewLogin(app pitaya.Pitaya) *Login {
	return &Login{
		app: app,
	}
}

func (m *Login) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("gate Login Login..., id = %d\n", in.GetRoleId())
	s := m.app.GetSessionFromCtx(ctx)
	// TODO: check roleId
	err := s.Bind(ctx, strconv.FormatInt(in.RoleId, 10))
	if err != nil {
		return nil, pitaya.Error(err, "RH-000", map[string]string{"failed": "bind"})
	}
	return &pb.LoginResponse{RoleId: in.GetRoleId()}, nil
}
