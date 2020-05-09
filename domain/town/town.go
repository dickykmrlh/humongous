package town

import (
	"context"
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
	State      string
}

type Politican struct {
	Name  string
	Party string
}

type TownCollection struct {
	collection *mongo.Collection
	ctx        context.Context
}

var townCollection *TownCollection

func GetInstance(ctx context.Context, db *mongo.Database) *TownCollection {
	if townCollection == nil {
		collection := db.Collection("towns")
		townCollection = &TownCollection{collection: collection, ctx: ctx}
	}
	return townCollection
}

func (t *TownCollection) InsertMany(towns []Town) ([]string, error) {

	var townsInterface []interface{}
	for _, t := range towns {
		townsInterface = append(townsInterface, t)
	}

	result, err := t.collection.InsertMany(t.ctx, townsInterface)
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
	result, err := t.collection.InsertOne(t.ctx, town)
	if err != nil {
		return "", err
	}
	objectID, _ := result.InsertedID.(primitive.ObjectID)
	return objectID.Hex(), nil
}

func (t *TownCollection) Find(opt *options.FindOptions) ([]Town, error) {
	cur, err := t.collection.Find(t.ctx, bson.M{}, opt)
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
	err := t.collection.FindOne(t.ctx, filter, opt).Decode(&result)
	if err != nil {
		return Town{}, err
	}
	return result, nil
}

func (t *TownCollection) UpdateOne(filter interface{}, update interface{}) error {
	_, err := t.collection.UpdateOne(t.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
