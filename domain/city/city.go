package city

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CitiesCollection struct {
	collection *mongo.Collection
	ctx        context.Context
}

var citiesCollection *CitiesCollection

func GetInstance(ctx context.Context, db *mongo.Database) *CitiesCollection {
	if citiesCollection == nil {
		collection := db.Collection("cities")
		citiesCollection = &CitiesCollection{collection: collection, ctx: ctx}
	}
	return citiesCollection
}

func (c *CitiesCollection) Aggregate(pipeline ...bson.D) ([]bson.M, error, ) {
	curr, err := c.collection.Aggregate(c.ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = curr.All(c.ctx, &results); err != nil {
		return nil,  err
	}

	return results, nil
}
