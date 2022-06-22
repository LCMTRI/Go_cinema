package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func Connect() {
	var URI = "mongodb+srv://lecaominhtri0701:lecaominhtri@cluster0.x7apzya.mongodb.net/?retryWrites=true&w=majority"
	fmt.Println("Starting the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}
	DB = client
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		fmt.Println("Disconnected")
	// 		panic(err)
	// 	}
	// }()
}
