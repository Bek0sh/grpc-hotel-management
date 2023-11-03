package service

import (
	"context"
	"fmt"
	"github.com/Bek0sh/hotel-management-booking/internal/models"
	"github.com/Bek0sh/hotel-management-booking/pkg/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"time"
)

const (
	constTime = "2006-01-02 15:04"
)

func (b *bookingService) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	id := primitive.NewObjectID()
	bookingId := id.Hex()

	roomType := struct {
		Type  string `bson:"type" json:"type"`
		Price int    `bson:"price" json:"price"`
	}{
		Type:  req.GetRoom().GetRoomType().GetRoomType(),
		Price: int(req.GetRoom().GetRoomType().GetPrice()),
	}

	room := struct {
		RoomNumber int `bson:"room_number" json:"room_number"`
		RoomType   struct {
			Type  string `bson:"type" json:"type"`
			Price int    `bson:"price" json:"price"`
		} `bson:"room_type" json:"room_type"`
	}{
		RoomNumber: int(req.GetRoom().GetRoomNumber()),
		RoomType:   roomType,
	}

	createBooking := models.Booking{
		Id:           id,
		BookingId:    bookingId,
		CheckInDate:  req.GetCheckInDate(),
		CheckOutDate: req.GetCheckOutDate(),
		Room:         room,
	}

	totalPrice, err := calculatePrice(req.GetCheckInDate(), req.GetCheckOutDate(), roomType.Price)
	if err != nil {
		b.log.Errorf("failed to create booking, invalid date, error: %v", err)
		return nil, err
	}
	createBooking.TotalPrice = totalPrice

	err = b.repo.CreateBooking(ctx, &createBooking)

	if err != nil {
		return nil, err
	}

	return &pb.CreateBookingResponse{TotalPrice: int32(totalPrice), Id: bookingId}, nil
}

func calculatePrice(in string, out string, price int) (int, error) {
	checkIn, err := time.Parse(constTime, in)
	if err != nil {
		return 0, err
	}
	checkOut, err := time.Parse(constTime, out)
	if err != nil {
		return 0, err
	}

	stayDuration := checkOut.Sub(checkIn)
	if stayDuration.Hours() < 24 {
		return 0, fmt.Errorf("you can stay minimum 1 day, less hours are forbidden, hours: %f", stayDuration.Hours())
	}
	var minutes float64 = 24 * 60

	numDays := stayDuration.Minutes() / minutes
	var totalPrice = float64(price) * numDays

	return int(totalPrice), nil
}

func (b *bookingService) DeleteBooking(ctx context.Context, req *pb.DeleteBookingRequest) (*pb.Empty, error) {
	err := b.repo.DeleteBooking(ctx, req.GetBookingId())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (b *bookingService) GetBookingsForCustomer(ctx context.Context, req *pb.GetBookingsForCustomerRequest) (*pb.GetBookingForCustomerResponse, error) {
	var response []*pb.Booking

	res, err := b.repo.GetBookingForCustomer(ctx, req.GetCustomerId())

	if err != nil {
		return nil, err
	}

	for _, v := range res {
		booking := &pb.Booking{
			CheckOutDate: v.CheckOutDate,
			CheckInDate:  v.CheckInDate,
			TotalPrice:   int32(v.TotalPrice),
			Room: &pb.Room{
				RoomNumber: int32(v.Room.RoomNumber),
				RoomType: &pb.RoomType{
					RoomType: v.Room.RoomType.Type,
					Price:    int32(v.Room.RoomType.Price),
				},
			},
		}
		response = append(response, booking)
	}

	return &pb.GetBookingForCustomerResponse{Bookings: response}, nil
}

func (b *bookingService) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.UpdateBookingResponse, error) {
	var message string

	booking, err := b.repo.GetBookingByRoomNumber(ctx, int(req.GetRoom().GetRoomNumber()))
	if err != nil {
		return nil, err
	}

	update := models.Booking{
		BookingId:    booking.BookingId,
		CheckInDate:  req.GetCheckInDate(),
		CheckOutDate: req.GetCheckOutDate(),
	}

	totalPrice, err := calculatePrice(update.CheckInDate, update.CheckOutDate, int(req.GetRoom().GetRoomType().GetPrice()))

	if err != nil {
		return nil, err
	}
	update.TotalPrice = totalPrice

	err = b.repo.UpdateBooking(ctx, &update)
	if err != nil {
		return nil, err
	}

	res := totalPrice - booking.TotalPrice

	if res < 0 {
		message = fmt.Sprintf("we need to return you %d", int(math.Abs(float64(res))))
	} else {
		message = fmt.Sprintf("you need to pay %d more", res)
	}

	return &pb.UpdateBookingResponse{Message: message}, nil
}
func (b *bookingService) GetAllBookings(ctx context.Context, _ *pb.Empty) (*pb.GetAllBookingsResponse, error) {
	var response []*pb.Booking

	res, err := b.repo.GetAllBookings(ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range res {
		booking := &pb.Booking{
			CheckOutDate: v.CheckOutDate,
			CheckInDate:  v.CheckInDate,
			TotalPrice:   int32(v.TotalPrice),
			Room: &pb.Room{
				RoomNumber: int32(v.Room.RoomNumber),
				RoomType: &pb.RoomType{
					RoomType: v.Room.RoomType.Type,
					Price:    int32(v.Room.RoomType.Price),
				},
			},
			CustomerId: *v.CustomerId,
		}
		response = append(response, booking)
	}

	return &pb.GetAllBookingsResponse{Bookings: response}, nil
}
