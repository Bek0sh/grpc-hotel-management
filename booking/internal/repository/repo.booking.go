package repository

import (
	"context"
	"github.com/Bek0sh/hotel-management-booking/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (b *booking) CreateBooking(ctx context.Context, req *models.Booking) error {
	_, err := b.InsertOne(ctx, req)
	if err != nil {
		b.log.Errorf("failed to create booking, error: %v", err)
		return err
	}

	return nil
}

func (b *booking) GetBookingById(ctx context.Context, id string) (*models.Booking, error) {
	response := models.Booking{}

	filter := bson.M{
		"booking_id": id,
	}

	err := b.FindOne(ctx, filter).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response, nil

}

func (b *booking) DeleteBooking(ctx context.Context, id string) error {

	filter := bson.M{
		"booking_id": id,
	}

	_, err := b.DeleteOne(ctx, filter)

	if err != nil {
		b.log.Errorf("failed to delete booking with id=%s, error: %v", id, err)
		return err
	}

	return nil
}

func (b *booking) GetBookingForCustomer(ctx context.Context, custId string) ([]models.Booking, error) {
	var response []models.Booking
	filter := bson.M{
		"customer_id": custId,
	}

	res, err := b.Find(ctx, filter)
	if err != nil {
		b.log.Errorf("failed to find bookings, error: %v", err)
		return nil, err
	}

	for res.Next(ctx) {
		booking := models.Booking{}
		err := res.Decode(&booking)

		if err != nil {
			b.log.Errorf("failed to iterate through available rooms, error: %v", err)
			return nil, err
		}

		response = append(response, booking)
	}

	return response, nil
}

func (b *booking) GetAllBookings(ctx context.Context) ([]models.Booking, error) {
	var response []models.Booking
	filter := bson.M{}

	res, err := b.Find(ctx, filter)
	if err != nil {
		b.log.Errorf("failed to find bookings, error: %v", err)
		return nil, err
	}

	for res.Next(ctx) {
		booking := models.Booking{}
		err := res.Decode(&booking)

		if err != nil {
			b.log.Errorf("failed to iterate through available rooms, error: %v", err)
			return nil, err
		}

		response = append(response, booking)
	}

	return response, nil
}

func (b *booking) GetBookingByRoomNumber(ctx context.Context, roomNumber int) (*models.Booking, error) {
	filter := bson.M{
		"room_id": roomNumber,
	}

	response := &models.Booking{}

	err := b.FindOne(ctx, filter).Decode(response)
	if err != nil {
		b.log.Errorf("failed to find booking with room number = %d, error: %v", roomNumber, err)
		return nil, err
	}

	return response, nil
}

func (b *booking) UpdateBooking(ctx context.Context, req *models.Booking) error {
	filter := bson.M{
		"booking_id": req.BookingId,
	}
	update := bson.M{
		"$set": bson.M{
			"check_in_date":  req.CheckInDate,
			"check_out_date": req.CheckOutDate,
			"room_id":        req.Room.RoomNumber,
			"total_price":    req.TotalPrice,
		},
	}

	_, err := b.UpdateOne(ctx, filter, update)

	if err != nil {
		b.log.Errorf("failed to update booking, error: %v", err)
		return err
	}

	return nil
}
