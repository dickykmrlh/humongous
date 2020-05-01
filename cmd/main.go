package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dickymrlh/humongous/db"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := db.GetMongoClient(ctx)
	defer client.Disconnect(ctx)

	result, err := client.ListDatabaseNames(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
