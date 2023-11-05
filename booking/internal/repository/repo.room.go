package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Bek0sh/hotel-management-booking/internal/models"
)

func (r *room) CreateRoom(ctx context.Context, req *models.Room) error {

	_, err := r.InsertOne(ctx, req)

	if err != nil {
		r.log.Errorf("failed to create user, error: %v", err)
		return fmt.Errorf("failed to create user, error: %v", err)
	}

	return nil
}

func (r *room) GetRoomByNumber(ctx context.Context, roomNumber int) (*models.Room, error) {
	response := &models.Room{}

	filter := bson.M{
		"room_number": roomNumber,
	}

	err := r.FindOne(ctx, filter).Decode(response)

	if err != nil {
		r.log.Errorf("failed to find room with number=%d, error: %v", roomNumber, err)
		return nil, err
	}

	return response, nil

}

func (r *room) GetAllAvailableRooms(ctx context.Context) ([]models.Room, error) {
	filter := bson.M{
		"is_available": true,
	}

	var response []models.Room
	res, err := r.Find(ctx, filter)
	if err != nil {
		r.log.Errorf("failed to find available rooms, error: %v", err)
		return nil, err
	}

	for res.Next(ctx) {
		room := models.Room{}
		err := res.Decode(&room)

		if err != nil {
			r.log.Errorf("failed to iterate through available rooms, error: %v", err)
			return nil, err
		}

		response = append(response, room)
	}

	return response, nil
}

func (r *room) UpdateAvailableness(ctx context.Context, roomNumber int, isAvailable bool) error {
	filter := bson.M{
		"room_number": roomNumber,
	}

	update := bson.M{
		"$set": bson.M{
			"is_available": isAvailable,
		},
	}
	_, err := r.UpdateOne(ctx, filter, update)
	if err != nil {
		r.log.Errorf("failed to update availability of room number=%d, error: %v", roomNumber, err)
		return err
	}

	return nil
}

func (r *room) UpdateRoom(ctx context.Context, req *models.Room) error {
	filter := bson.M{
		"room_number": req.RoomNumber,
	}

	update := bson.M{
		"$set": bson.M{
			"room_type": bson.M{
				"type":  req.RoomType.Type,
				"price": req.RoomType.Price,
			},
		},
	}

	_, err := r.UpdateOne(ctx, filter, update)
	if err != nil {
		r.log.Errorf("failed to update availability of room number=%d, error: %v", req.RoomNumber, err)
		return err
	}

	return nil
}

func (r *room) DeleteRoomByNumber(ctx context.Context, roomNumber int) error {
	filter := bson.M{
		"room_number": roomNumber,
	}

	_, err := r.DeleteOne(ctx, filter)

	if err != nil {
		r.log.Errorf("failed to delete room with number = %d, error: %v", roomNumber, err)
		return err
	}

	return nil
}
