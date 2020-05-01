package town

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TownCollection struct {
	c *mongo.Collection
}

var townCollection *TownCollection

func GetInstance(ctx context.Context) *TownCollection {
	if townCollection == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		db := client.Database("world")
		collection := db.Collection("town")
		townCollection = &TownCollection{c: collection}
	}
	return townCollection
}
