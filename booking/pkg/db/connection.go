package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Bek0sh/hotel-management-booking/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

type Database struct {
	*mongo.Database
}

func GetDatabase() *Database {
	return &Database{db}
}

func GetInstance(cfg config.Config) {
	mongoUrl := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		cfg.Mongo.Username,
		cfg.Mongo.Password,
		cfg.Mongo.Host,
		cfg.Mongo.Port,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))

	if err != nil {
		panic(err)
	}

	db = cl.Database(cfg.Mongo.DbName)

	fmt.Println("connected to mongodb")
}
