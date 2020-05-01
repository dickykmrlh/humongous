package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dickymrlh/humongous/domain/town"
	t "github.com/dickymrlh/humongous/usecase/town"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c := town.GetInstance(ctx)

	town := t.Town{
		Name:       "New York",
		Population: 22200000,
		LastCensus: time.Date(2016, 7, 1, 0, 0, 0, 0, time.Local),
		FamousFor:  []string{"the MOMA", "food", "Derek Jeter"},
		Mayor:      t.Politican{Name: "Bill de Blasio", Party: "I"},
	}

	data, err := bson.Marshal(town)
	if err != nil {
		panic(err)
	}

	result, err := c.Mc.InsertOne(ctx, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.InsertedID)

}
