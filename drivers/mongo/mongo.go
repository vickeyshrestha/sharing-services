package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDriver(mongoURL string) (*mongo.Client, error) {
	if len(mongoURL) == 0 {
		return nil, errors.New("missing mongoURL")
	}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", mongoURL))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}
	return client, nil
}
