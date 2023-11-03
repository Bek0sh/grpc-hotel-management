package repository

import (
	"context"

	"github.com/Bek0sh/hotel-management-auth/internal/models"
	"github.com/Bek0sh/hotel-management-auth/pkg/logging"
	mongodb "github.com/Bek0sh/hotel-management-auth/pkg/mongoDB"
)

type IRepository interface {
	CreateUser(ctx context.Context, req *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type repo struct {
	logger logging.Logger
	c      *mongodb.Database
}

func New(l logging.Logger, col *mongodb.Database) IRepository {
	return &repo{logger: l, c: col}
}
