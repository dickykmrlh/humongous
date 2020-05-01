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

type Collection struct {
	m *mongo.Collection
}

var c *Collection

func GetInstance(ctx context.Context) *Collection {
	if c == nil {
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
		c = &Collection{m: collection}
	}
	return c
}

func (c *Collection) Insert(ctx context.Context, towns []Town) ([]string, error) {

	var townsInterface []interface{}
	for _, t := range towns {
		townsInterface = append(townsInterface, t)
	}

	result, err := c.m.InsertMany(ctx, townsInterface)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, id := range result.InsertedIDs {
		ids = append(ids, fmt.Sprintf("%v", id))
	}

	return ids, nil
}
