package city

import (
	"fmt"
	"github.com/dickymrlh/humongous/domain/city"
	"go.mongodb.org/mongo-driver/bson"
)

func PlayAroundWithTownAggregate(collection *city.CitiesCollection) {
	matchStage :=  bson.D{
		{"$match", bson.D{
			{"timezone", bson.D{{"$eq", "Europe/London"}} },
		}},
	}

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "averagePopulation"},
			{"avgPop", bson.D{ {"$avg", "$population"} }},
		}},
	}

	sortStage := bson.D {
		{"$sort", bson.D{{"population", -1}}},
	}

	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"name", 1},
			{"population", 1},
		}},
	}

	results, err := collection.Aggregate(matchStage, groupStage)
	if err != nil {
		panic(err)
	}

	for _, r := range results {
		fmt.Println(r)
	}

	results, err = collection.Aggregate(matchStage, sortStage, projectStage)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(results[i])
	}
}