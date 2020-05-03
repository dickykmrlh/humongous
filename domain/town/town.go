package town

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	c   *mongo.Collection
	ctx context.Context
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
		townCollection = &TownCollection{c: collection, ctx: ctx}
	}
	return townCollection
}

func (t *TownCollection) InsertMany(towns []Town) ([]string, error) {

	var townsInterface []interface{}
	for _, t := range towns {
		townsInterface = append(townsInterface, t)
	}

	result, err := t.c.InsertMany(t.ctx, townsInterface)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, id := range result.InsertedIDs {
		objectID, _ := id.(primitive.ObjectID)
		ids = append(ids, objectID.Hex())
	}

	return ids, nil
}

func (t *TownCollection) InsertOne(town Town) (string, error) {
	result, err := t.c.InsertOne(t.ctx, town)
	if err != nil {
		return "", err
	}
	objectID, _ := result.InsertedID.(primitive.ObjectID)
	return objectID.Hex(), nil
}

func (t *TownCollection) Find(opt *options.FindOptions) ([]Town, error) {
	cur, err := t.c.Find(t.ctx, bson.M{}, opt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(t.ctx)

	var results []Town
	for cur.Next(t.ctx) {

		var elem Town
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (t *TownCollection) FindOne(filter interface{}, opt *options.FindOneOptions) (Town, error) {
	var result Town
	err := t.c.FindOne(t.ctx, filter, opt).Decode(&result)
	if err != nil {
		return Town{}, err
	}
	return result, nil
}
