package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	UserCollection = "users"
)

// connect to mongodb
func InitDB(addr string, dbname string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))
	if err != nil {
		return nil, err
	}

	// test ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database(dbname)
	// ensure username in the index
	col := db.Collection(UserCollection)
	mod := mongo.IndexModel{
		Keys: bson.M{
			"username": 1, // index in ascending order
		},
		Options: options.Index().SetUnique(true),
	}
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	if _, err = col.Indexes().CreateOne(ctx, mod, opts); err != nil {
		return nil, err
	}

	return client.Database(dbname), nil
}
