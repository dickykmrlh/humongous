package town

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Town struct {
	Name       string
	Population int64
	LastCensus time.Time
	FamousFor  []string
	Mayor      Politican
}

type Politican struct {
	Name  string
	Party string
}

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
		collection := db.Collection("towns")
		townCollection = &TownCollection{c: collection}
	}
	return townCollection
}

func (t *TownCollection) Insert(ctx context.Context, towns []Town) ([]string, error) {

	if len(towns) > 1 {
		var townsInterface []interface{}
		for _, t := range towns {
			townsInterface = append(townsInterface, t)
		}

		resultMany, err := t.c.InsertMany(ctx, townsInterface)
		if err != nil {
			return nil, err
		}

		var ids []string
		for _, id := range resultMany.InsertedIDs {
			ids = append(ids, fmt.Sprintf("%v", id))
		}

		return ids, nil
	}

	resultOne, err := t.c.InsertOne(ctx, towns[0])
	if err != nil {
		return nil, err
	}
	return []string{fmt.Sprintf("%v", resultOne.InsertedID)}, nil
}

//func (t *TownCollection) Find()
