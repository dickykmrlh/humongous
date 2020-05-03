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

func (t *TownCollection) InsertMany(ctx context.Context, towns []Town) ([]string, error) {

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

func (t *TownCollection) InsertOne(ctx context.Context, town Town) (string, error) {
	resultOne, err := t.c.InsertOne(ctx, town)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", resultOne.InsertedID), nil
}

/*
func (t *TownCollection) Find() []Town{
	findOptions := options.Find()
findOptions.SetLimit(2)

var results []Town

// Passing bson.D{{}} as the filter matches all documents in the collection
cur, err := t.c.Find(context.TODO(), bson.D{{}}, findOptions)
if err != nil {
    log.Fatal(err)
}

// Finding multiple documents returns a cursor
// Iterating through the cursor allows us to decode documents one at a time
for cur.Next(context.TODO()) {

    // create a value into which the single document can be decoded
    var elem Town
    err := cur.Decode(&elem)
    if err != nil {
        log.Fatal(err)
    }

    results = append(results, elem)
}

if err := cur.Err(); err != nil {
    log.Fatal(err)
}

// Close the cursor once finished
cur.Close(context.TODO())

fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}
*/
