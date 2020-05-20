package main

import (
	"context"
	"github.com/dickymrlh/humongous/domain/city"
	city2 "github.com/dickymrlh/humongous/usecase/city"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("world")

	collection := city.GetInstance(ctx, db)
	city2.PlayAroundWithTownAggregate(collection)
}
