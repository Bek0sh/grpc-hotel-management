package main

import (
	"github.com/Bek0sh/hotel-management-booking/internal/config"
	"github.com/Bek0sh/hotel-management-booking/internal/repository"
	"github.com/Bek0sh/hotel-management-booking/internal/service"
	"github.com/Bek0sh/hotel-management-booking/pkg/db"
	"github.com/Bek0sh/hotel-management-booking/pkg/logging"
	"github.com/Bek0sh/hotel-management-booking/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	cfg := config.GetConfig()
	logging.Init()
	logger := logging.GetLogger()

	db.GetInstance(*cfg)
	database := db.GetDatabase()

	roomTypeRepo := repository.NewRoomTypeRepository(logger, database.Database)
	roomRepo := repository.NewRoomRepo(logger, database.Database)
	bookingRepo := repository.NewBookingRepo(logger, database.Database)

	roomTypeSrv := service.NewRoomTypeService(logger, roomTypeRepo)
	roomSrv := service.NewRoomService(logger, roomRepo)
	bookingSrv := service.NewBookingService(logger, bookingRepo)

	listener, err := net.Listen("tcp", cfg.Run.Port)

	if err != nil {
		logger.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterRoomTypeServiceServer(grpcServer, roomTypeSrv)
	pb.RegisterRoomServiceServer(grpcServer, roomSrv)
	pb.RegisterBookingServiceServer(grpcServer, bookingSrv)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal(err)
	}
}
