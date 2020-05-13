package phone

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type PhoneCollection struct {
	collection *mongo.Collection
	ctx        context.Context
}

type Phone struct {
	ID         int64     `bson:"_id"`
	Components Component `bson:"components"`
	Display    string    `bson:"display"`
}

type Component struct {
	Country int `bson:"country"`
	Area    int `bson:"area"`
	Prefix  int `bson:"prefix"`
	Number  int `bson:"number"`
}

var phoneCollection *PhoneCollection

func GetInstance(ctx context.Context, db *mongo.Database) *PhoneCollection {
	if phoneCollection == nil {
		collection := db.Collection("phones")
		phoneCollection = &PhoneCollection{collection: collection, ctx: ctx}
	}
	return phoneCollection
}

func (p *PhoneCollection) InsertOne(phone Phone) error {
	_, err := p.collection.InsertOne(p.ctx, phone)
	if err != nil {
		return err
	}
	return nil
}
