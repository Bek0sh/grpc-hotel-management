package main

import (
	"github.com/Bek0sh/hotel-management-auth/internal/config"
	"github.com/Bek0sh/hotel-management-auth/internal/repository"
	"github.com/Bek0sh/hotel-management-auth/internal/service"
	"github.com/Bek0sh/hotel-management-auth/pkg/logging"
	mongodb "github.com/Bek0sh/hotel-management-auth/pkg/mongoDB"
	"github.com/Bek0sh/hotel-management-auth/pkg/pb"
	"github.com/Bek0sh/hotel-management-auth/pkg/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var cfg *config.Config
var srv pb.AuthServiceServer

func init() {
	var err error

	logging.Init()
	logger := logging.GetLogger()
	cfg = config.GetConfig()
	jwt, err := token.NewJwtMaker(cfg.JWT.SecretKey)
	if err != nil {
		logger.Fatal(err)
	}

	mongodb.DBInstance(cfg)
	db := mongodb.GetCollection()
	repo := repository.New(logger, db)
	srv = service.NewService(repo, logger, jwt, *cfg)
}

func main() {

	listener, err := net.Listen("tcp", cfg.Run.Port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, srv)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
