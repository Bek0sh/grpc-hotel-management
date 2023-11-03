package service

import (
	"github.com/Bek0sh/hotel-management-booking/internal/repository"
	"github.com/Bek0sh/hotel-management-booking/pkg/logging"
	"github.com/Bek0sh/hotel-management-booking/pkg/pb"
)

type bookingService struct {
	log  *logging.Logger
	repo repository.BookingRepository
	pb.UnimplementedBookingServiceServer
}

type roomService struct {
	log  *logging.Logger
	repo repository.RoomRepository
	pb.UnimplementedRoomServiceServer
}

type roomTypeService struct {
	log  *logging.Logger
	repo repository.RoomTypeRepository
	pb.UnimplementedRoomTypeServiceServer
}

func NewRoomService(log *logging.Logger, repo repository.RoomRepository) pb.RoomServiceServer {
	return &roomService{log: log, repo: repo}
}

func NewRoomTypeService(log *logging.Logger, repo repository.RoomTypeRepository) pb.RoomTypeServiceServer {
	return &roomTypeService{
		log:  log,
		repo: repo,
	}
}

func NewBookingService(log *logging.Logger, repo repository.BookingRepository) pb.BookingServiceServer {
	return &bookingService{log: log, repo: repo}
}
