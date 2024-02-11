package user

import (
	desc "github.com/terdenan/msc_auth/pkg/user_v1"
)

type UserServer struct {
	desc.UnimplementedUserV1Server
}

func NewUserServer() *UserServer {
	return &UserServer{}
}
