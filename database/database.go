package database

import (
	"context"
	"fmt"

	"github.com/waliqueiroz/podcastr-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/", config.DBHost, config.DBPort))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(config.DBDatabase), nil
}
