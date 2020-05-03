package town

import (
	"fmt"
	"sync"
	"time"

	"github.com/dickymrlh/humongous/domain/town"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PlayAroundWithTown(townCollection *town.TownCollection) {
	fmt.Println("###################################################################################")
	ids, err := InsertTown(townCollection)
	if err != nil {
		panic(err)
	}

	opt := options.Find()
	// find all
	towns, err := townCollection.Find(opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(towns)
	fmt.Println("=====================================")

	// find with limit and sort
	opt.SetLimit(2)
	opt.SetSort(bson.D{{"population", 1}})
	towns, err = townCollection.Find(opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(towns)
	fmt.Println("=====================================")

	if ids == nil {
		return
	}

	// find One with object ID
	objID, err := primitive.ObjectIDFromHex(ids[1])
	if err != nil {
		panic(err)
	}

	findOne := options.FindOne()
	town, err := townCollection.FindOne(bson.D{{"_id", objID}}, findOne)
	if err != nil {
		panic(err)
	}
	fmt.Println(town)

	// find One with object ID and return Name of the city only
	findOne.SetProjection(bson.D{{"name", 1}})
	town, err = townCollection.FindOne(bson.D{{"_id", objID}}, findOne)
	if err != nil {
		panic(err)
	}

	// find One With population in range
	// $lt = less than
	// $gt = greater than
	findOne.SetProjection(bson.D{{"name", 1}, {"population", 1}})
	town, err = townCollection.FindOne(bson.M{
		"population": bson.M{
			"$lt": 1000000,
			"$gt": 10000,
		},
	}, findOne)
	if err != nil {
		panic(err)
	}

	// find One matching partial values using regEx
	// options "i" = for case incensitive
	regex := bson.M{"$regex": primitive.Regex{Pattern: "moma", Options: "i"}}
	findOne.SetProjection(bson.D{{"name", 1}, {"famousfor", 1}})
	town, err = townCollection.FindOne(bson.M{"famousfor": regex}, findOne)
	if err != nil {
		panic(err)
	}
	fmt.Println(town)
	fmt.Println()
	fmt.Println("###################################################################################")
}

func InsertTown(townCollection *town.TownCollection) (ids []string, err error) {

	if isDocumentAlreadyExist(townCollection) {
		return
	}

	towns := []town.Town{
		town.Town{
			Name:       "New York",
			Population: 22200000,
			LastCensus: time.Date(2016, 7, 1, 0, 0, 0, 0, time.Local),
			FamousFor:  []string{"the MOMA", "food", "Derek Jeter"},
			Mayor:      town.Politican{Name: "Bill de Blasio", Party: "I"},
		},
		town.Town{
			Name:       "Punxsutawney",
			Population: 6200,
			LastCensus: time.Date(2016, 1, 31, 0, 0, 0, 0, time.Local),
			FamousFor:  []string{"Punxsutawney Phil"},
			Mayor:      town.Politican{Name: "Richard Alexander"},
		},
		town.Town{
			Name:       "Portland",
			Population: 582000,
			LastCensus: time.Date(2016, 9, 20, 0, 0, 0, 0, time.Local),
			FamousFor:  []string{"berr", "food", "Portlandia"},
			Mayor:      town.Politican{Name: "Ted Wheeler", Party: "D"},
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ids, err = townCollection.InsertMany(towns)
	}()
	wg.Wait()
	if err != nil {
		return
	}

	return
}

func isDocumentAlreadyExist(townCollection *town.TownCollection) bool {
	towns, find_err := townCollection.Find(options.Find())
	if find_err != nil {
		panic(find_err)
	}
	if towns != nil {
		return true
	}
	return false
}
