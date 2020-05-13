package main

import (
	"context"
	"log"
	"time"

	//"github.com/dickymrlh/humongous/domain/town"

	//t "github.com/dickymrlh/humongous/usecase/town"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/dickymrlh/humongous/domain/phone"
	p "github.com/dickymrlh/humongous/usecase/phone"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("world")

	phoneCollection := phone.GetInstance(ctx, db)
	p.PlayAroundWithPhone(phoneCollection)
	//townCollection := town.GetInstance(ctx, db)
	//t.PlayAroundWithTown(townCollection)
	//countriesCollection := country.GetInstance(ctx, db)
	//c.PlayAroundWithCountry(countriesCollection)
}
