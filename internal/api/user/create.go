package user

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	desc "github.com/terdenan/msc_auth/pkg/user_v1"
)

func (s *UserServer) Create(_ context.Context, _ *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}
