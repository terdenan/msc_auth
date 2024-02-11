package user

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/terdenan/msc_auth/pkg/user_v1"
)

func (s *UserServer) Get(_ context.Context, _ *desc.GetRequest) (*desc.GetResponse, error) {
	return &desc.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_user,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}
