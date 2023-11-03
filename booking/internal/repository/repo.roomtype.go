package repository

import (
	"context"
	"fmt"
	"github.com/Bek0sh/hotel-management-booking/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *roomType) CreateRoomType(ctx context.Context, roomType *models.RoomType) error {

	if ok := r.CategoryExists(ctx, roomType.Type); ok {
		r.log.Errorf("failed to create room type, error: This room type exists")
		return fmt.Errorf("failed to create room type, error: This room type exists")
	}

	_, err := r.InsertOne(ctx, roomType)

	if err != nil {
		r.log.Errorf("failed to create room type, error: %v", err)
		return err
	}

	return nil
}

func (r *roomType) CategoryExists(ctx context.Context, category string) bool {
	filter := bson.M{
		"type": category,
	}
	res, err := r.CountDocuments(ctx, filter)
	if err != nil {
		r.log.Error("failed to count categories")
		return true
	}

	if res != 0 {
		r.log.Error("category type must be unique")
		return true
	}

	return false
}
func (r *roomType) GetRoomTypeByType(ctx context.Context, id string) (*models.RoomType, error) {
	response := &models.RoomType{}

	filter := bson.M{
		"type": id,
	}

	err := r.FindOne(ctx, filter).Decode(response)

	if err != nil {
		r.log.Errorf("failed to find room type with id: %s, error: %v", id, err)
		return nil, err
	}

	return response, nil
}
func (r *roomType) DeleteRoomTypeByType(ctx context.Context, id string) error {
	filter := bson.M{
		"type": id,
	}

	_, err := r.DeleteOne(ctx, filter)
	if err != nil {
		r.log.Errorf("failed to delete room type with id = %s", id)
		return err
	}

	return nil
}

func (r *roomType) GetAllRoomTypes(ctx context.Context) ([]models.RoomType, error) {
	filter := bson.M{}

	res, err := r.Find(ctx, filter)
	var response []models.RoomType
	if err != nil {
		r.log.Errorf("failed to find available rooms, error: %v", err)
		return nil, err
	}

	for res.Next(ctx) {
		roomT := models.RoomType{}
		err := res.Decode(&roomT)

		if err != nil {
			r.log.Errorf("failed to iterate through available rooms, error: %v", err)
			return nil, err
		}

		response = append(response, roomT)
	}

	return response, nil
}

func (r *roomType) UpdateRoomType(ctx context.Context, id string, price int) error {
	filter := bson.M{
		"room_type_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"price": price,
		},
	}

	_, err := r.UpdateOne(ctx, filter, update)

	if err != nil {
		r.log.Errorf("failed to update room type, error: %v", err)
		return err
	}

	return nil
}
