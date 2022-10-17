package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongorm_client struct {
	client *mongo.Client
}

var (
	client *mongo.Client
)

func Connect(connectionString string) (*mongorm_client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	return &mongorm_client{client}, nil
}
