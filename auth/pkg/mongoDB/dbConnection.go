package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/Bek0sh/hotel-management-auth/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Collection

type Database struct {
	*mongo.Collection
}

func GetCollection() *Database {
	return &Database{client}
}

func DBInstance(cfg *config.Config) {
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

	db := cl.Database(cfg.Mongo.DbName)

	client = db.Collection(cfg.Mongo.CollectionName)

	fmt.Println("connected to mongodb")
}
