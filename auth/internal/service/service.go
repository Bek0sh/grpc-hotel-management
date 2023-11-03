package service

import (
	"github.com/Bek0sh/hotel-management-auth/pkg/pb"

	"github.com/Bek0sh/hotel-management-auth/internal/config"
	"github.com/Bek0sh/hotel-management-auth/internal/repository"
	"github.com/Bek0sh/hotel-management-auth/pkg/logging"
	"github.com/Bek0sh/hotel-management-auth/pkg/token"
)

type service struct {
	repo   repository.IRepository
	logger logging.Logger
	jwt    token.Maker
	cfg    config.Config
	pb.UnimplementedAuthServiceServer
}

func NewService(repo repository.IRepository, logger logging.Logger, jwt token.Maker, cfg config.Config) pb.AuthServiceServer {
	return &service{repo: repo, logger: logger, jwt: jwt, cfg: cfg}
}
