package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/terdenan/msc_auth/pkg/user_v1"
)

func (s *UserServer) Update(_ context.Context, _ *desc.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
