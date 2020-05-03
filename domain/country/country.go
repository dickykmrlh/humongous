package country

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Country struct {
	ID      string   `bson:"_id"`
	Name    string   `bson:"name"`
	Exports []Export `bson:"exports"`
}

type Export struct {
	Name      string
	Tasty     bool
	Condiment bool
}

type CountriesCollection struct {
	collection *mongo.Collection
	ctx        context.Context
}

var countriesCollection *CountriesCollection

func GetInstance(ctx context.Context, db *mongo.Database) *CountriesCollection {
	if countriesCollection == nil {
		collection := db.Collection("countries")
		countriesCollection = &CountriesCollection{collection: collection, ctx: ctx}
	}
	return countriesCollection
}

func (c *CountriesCollection) InsertMany(countries []Country) error {

	var townsInterface []interface{}
	for _, t := range countries {
		townsInterface = append(townsInterface, t)
	}

	_, err := c.collection.InsertMany(c.ctx, townsInterface)
	if err != nil {
		return err
	}

	return nil
}

func (c *CountriesCollection) InsertOne(country Country) error {
	_, err := c.collection.InsertOne(c.ctx, country)
	if err != nil {
		return err
	}
	return nil
}

func (c *CountriesCollection) Find(opt *options.FindOptions) ([]Country, error) {
	cur, err := c.collection.Find(c.ctx, bson.M{}, opt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(c.ctx)

	var results []Country
	cur.All(c.ctx, &results)
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
