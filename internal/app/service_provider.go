package app

import (
	"log"

	"github.com/terdenan/msc_auth/internal/api/user"
	"github.com/terdenan/msc_auth/internal/config"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	userServer *user.UserServer
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) UserServer() *user.UserServer {
	if s.userServer == nil {
		s.userServer = user.NewUserServer()
	}

	return s.userServer
}
