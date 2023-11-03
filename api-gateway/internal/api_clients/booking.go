package api_clients

import (
	"context"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"time"
)

type BookingClient struct {
	cl  pb.BookingServiceClient
	log *logging.Logger
}

func NewBookingClient(conn *grpc.ClientConn, log *logging.Logger) *BookingClient {
	client := pb.NewBookingServiceClient(conn)
	return &BookingClient{cl: client, log: log}
}

func (b *BookingClient) CreateBooking(booking *models.Booking) (string, int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.CreateBookingRequest{
		Room:         booking.Room.ToProto(),
		CustomerId:   booking.CustomerId,
		CheckInDate:  booking.CheckInDate,
		CheckOutDate: booking.CheckOutDate,
	}

	res, err := b.cl.CreateBooking(ctx, req)
	if err != nil {
		b.log.Errorf("failed to create booking, error: %v", err)
		return "", 0, err
	}

	return res.GetId(), res.GetTotalPrice(), nil
}

func (b *BookingClient) GetAllBookings() ([]models.Booking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := b.cl.GetAllBookings(ctx, &pb.Empty{})
	if err != nil {
		b.log.Errorf("failed to get all bookings, error: %v", err)
		return nil, err
	}

	var response []models.Booking

	for _, v := range res.GetBookings() {
		booking := models.Booking{
			CheckOutDate: v.GetCheckOutDate(),
			CheckInDate:  v.GetCheckInDate(),
			TotalPrice:   v.GetTotalPrice(),
			Room: models.Room{
				RoomNumber: v.GetRoom().GetRoomNumber(),
				RoomType: models.RoomType{
					Type:  v.GetRoom().GetRoomType().GetRoomType(),
					Price: v.GetRoom().GetRoomType().GetPrice(),
				},
			},
		}

		response = append(response, booking)
	}

	return response, nil
}

func (b *BookingClient) GetCustomersBookings(custid string) ([]models.Booking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.GetBookingsForCustomerRequest{CustomerId: custid}

	res, err := b.cl.GetBookingsForCustomer(ctx, req)
	if err != nil {
		b.log.Errorf("failed to get bookings for customer = %s, error: %v", custid, err)
		return nil, err
	}

	var response []models.Booking

	for _, v := range res.GetBookings() {
		booking := models.Booking{
			CheckOutDate: v.GetCheckOutDate(),
			CheckInDate:  v.GetCheckInDate(),
			TotalPrice:   v.GetTotalPrice(),
			Room: models.Room{
				RoomNumber: v.GetRoom().GetRoomNumber(),
				RoomType: models.RoomType{
					Type:  v.GetRoom().GetRoomType().GetRoomType(),
					Price: v.GetRoom().GetRoomType().GetPrice(),
				},
			},
		}

		response = append(response, booking)
	}

	return response, nil
}

func (b *BookingClient) DeleteBooking(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.DeleteBookingRequest{BookingId: id}

	_, err := b.cl.DeleteBooking(ctx, req)
	if err != nil {
		b.log.Errorf("failed to delete booking with id=%s, error: %v", id, err)
		return err
	}

	return nil
}

func (b *BookingClient) UpdateBooking(booking *models.Booking) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.UpdateBookingRequest{
		Room:         booking.Room.ToProto(),
		CheckInDate:  booking.CheckInDate,
		CheckOutDate: booking.CheckOutDate,
	}

	res, err := b.cl.UpdateBooking(ctx, req)
	if err != nil {
		b.log.Errorf("failed to update booking, error: %v", err)
		return "", err
	}

	return res.GetMessage(), nil
}
