package databse

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc

func Close() {

	defer cancel()

	defer func() {
		if err := Client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func Connect(uri string) {

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Database")
}
