package helpers

import (
	"context"
	"time"

	"github.com/randy1burrell/toggle-game/pkg/application"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetClient gets an instance of a mongodb client, collection and context
func GetClient() (*mongo.Client, *mongo.Collection, *context.Context, error) {
	config := application.Get()
	db  := config.DB
	client, err := mongo.NewClient(options.Client().ApplyURI(db.GetDBConnStr()))
	if err !=  nil {
		return nil, nil, nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return nil, nil, nil, err
	}

	dbsName := client.Database(db.DbName)
	coll := dbsName.Collection(db.Collection)

	return client, coll, &ctx, err
}
