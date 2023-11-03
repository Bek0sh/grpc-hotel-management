package repository

import (
	"context"
	"errors"

	"github.com/Bek0sh/hotel-management-auth/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) CreateUser(ctx context.Context, req *models.User) error {

	if r.emailExists(ctx, req.Email) {
		return errors.New("this email is already exits")
	}

	_, err := r.c.Collection.InsertOne(ctx, req)
	if err != nil {
		r.logger.Error("failed to create user")
		return err
	}
	return nil
}

func (r *repo) emailExists(ctx context.Context, email string) bool {
	filter := bson.M{
		"email": email,
	}

	count, err := r.c.Collection.CountDocuments(ctx, filter)
	if err != nil {
		r.logger.Errorf("failed to count number of emails, error: %v", err)
		return true
	}
	if count != 0 {
		r.logger.Error("user provided existing email")
		return true
	}

	return false
}

func (r *repo) GetUserById(ctx context.Context, id string) (*models.User, error) {
	var resp models.User
	filter := bson.M{
		"user_id": id,
	}

	err := r.c.Collection.FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		r.logger.Errorf("failed to find user with user_id=%s", id)
		return nil, err
	}

	return &resp, nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var resp models.User
	filter := bson.M{
		"email": email,
	}

	err := r.c.Collection.FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		r.logger.Errorf("failed to find user with email=%s", email)
		return nil, err
	}

	return &resp, nil
}

func (r *repo) DeleteUser(ctx context.Context, id string) error {
	filter := bson.M{
		"user_id": id,
	}

	_, err := r.c.Collection.DeleteOne(ctx, filter)
	if err != nil {
		r.logger.Errorf("failed to delete user with user_id=%s", id)
		return err
	}

	return nil
}

func (r *repo) UpdateUser(ctx context.Context, id string) (*models.User, error) {
	return &models.User{}, nil
}
