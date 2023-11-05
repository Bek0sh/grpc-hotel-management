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
	UpdateRoom(ctx context.Context, req *models.Room) error
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
}

type booking struct {
	log *logging.Logger
	*mongo.Collection
}

func newRoomRepo(log *logging.Logger, db *mongo.Collection) RoomRepository {

	return &room{Collection: db, log: log}
}

func newBookingRepo(log *logging.Logger, db *mongo.Collection) BookingRepository {

	return &booking{Collection: db, log: log}
}

func newRoomTypeRepository(log *logging.Logger, db *mongo.Collection) RoomTypeRepository {
	return &roomType{Collection: db, log: log}
}

type Repository struct {
	Room     RoomRepository
	Booking  BookingRepository
	RoomType RoomTypeRepository
}

func NewRepository(log *logging.Logger, db *mongo.Database) *Repository {
	collectionRoomType := db.Collection(roomTypeCollection)
	collectionRoom := db.Collection(roomCollection)
	collectionBooking := db.Collection(bookingCollection)

	roomTypeRepo := newRoomTypeRepository(log, collectionRoomType)
	roomRepo := newRoomRepo(log, collectionRoom)
	bookingRepo := newBookingRepo(log, collectionBooking)

	return &Repository{Room: roomRepo, RoomType: roomTypeRepo, Booking: bookingRepo}
}
