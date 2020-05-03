package main

import (
	"context"
	"log"
	"time"

	"github.com/dickymrlh/humongous/domain/country"
	c "github.com/dickymrlh/humongous/usecase/country"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("world")

	//townCollection := town.GetInstance(ctx, db)
	//t.PlayAroundWithTown(townCollection)
	countriesCollection := country.GetInstance(ctx, db)
	c.PlayAroundWithCountry(countriesCollection)
}
