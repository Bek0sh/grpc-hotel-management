package repository

import (
	"context"
	"github.com/Bek0sh/hotel-management-booking/pkg/logging"

	"github.com/Bek0sh/hotel-management-booking/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	roomTypeCollection = "roomType"
	roomCollection     = "room"
	bookingCollection  = "booking"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, req *models.Booking) error
	DeleteBooking(ctx context.Context, id string) error
	GetAllBookings(ctx context.Context) ([]models.Booking, error)
	GetBookingForCustomer(ctx context.Context, custId string) ([]models.Booking, error)
	GetBookingByRoomNumber(ctx context.Context, roomNumber int) (*models.Booking, error)
	UpdateBooking(ctx context.Context, req *models.Booking) error
}

type RoomRepository interface {
	CreateRoom(ctx context.Context, req *models.Room) error
	GetRoomByNumber(ctx context.Context, roomNumber int) (*models.Room, error)
	GetAllAvailableRooms(ctx context.Context) ([]models.Room, error)
	DeleteRoomByNumber(ctx context.Context, roomNumber int) error
	UpdateRoom(ctx context.Context, req *models.Room, isAvailable bool) error
	UpdateAvailableness(ctx context.Context, roomNumber int, isAvailable bool) error
}

type RoomTypeRepository interface {
	CreateRoomType(ctx context.Context, roomType *models.RoomType) error
	GetRoomTypeByType(ctx context.Context, id string) (*models.RoomType, error)
	DeleteRoomTypeByType(ctx context.Context, id string) error
	UpdateRoomType(ctx context.Context, id string, price int) error
	CategoryExists(ctx context.Context, category string) bool
	GetAllRoomTypes(ctx context.Context) ([]models.RoomType, error)
}

type roomType struct {
	log *logging.Logger
	*mongo.Collection
}

type room struct {
	log *logging.Logger
	*mongo.Collection
	repo RoomTypeRepository
}

type booking struct {
	log *logging.Logger
	*mongo.Collection
	repo RoomRepository
}

func NewRoomRepo(log *logging.Logger, db *mongo.Database) RoomRepository {
	repo := NewRoomTypeRepository(log, db)
	collection := db.Collection(roomCollection)
	return &room{Collection: collection, repo: repo, log: log}
}

func NewBookingRepo(log *logging.Logger, db *mongo.Database) BookingRepository {
	repo := NewRoomRepo(log, db)
	collection := db.Collection(bookingCollection)
	return &booking{Collection: collection, repo: repo, log: log}
}

func NewRoomTypeRepository(log *logging.Logger, db *mongo.Database) RoomTypeRepository {
	collection := db.Collection(roomTypeCollection)
	return &roomType{Collection: collection, log: log}
}
