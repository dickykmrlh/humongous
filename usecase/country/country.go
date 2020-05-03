package country

import (
	"fmt"
	"sync"

	"github.com/dickymrlh/humongous/domain/country"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PlayAroundWithCountry(countriesCollection *country.CountriesCollection) {
	fmt.Println("###################################################################################")
	ids, err := InsertTown(countriesCollection)
	if err != nil {
		panic(err)
	}
	fmt.Println(ids)
	/*
		opt := options.Find()
		// find all
		towns, err := countriesCollection.Find(opt)
		if err != nil {
			panic(err)
		}
		fmt.Println(towns)
		fmt.Println("=====================================")

		// find with limit and sort
		opt.SetLimit(2)
		opt.SetSort(bson.D{{"population", 1}})
		towns, err = countriesCollection.Find(opt)
		if err != nil {
			panic(err)
		}
		fmt.Println(towns)
		fmt.Println("=====================================")

		// find One with object ID
		objID, err := primitive.ObjectIDFromHex(ids[1])
		if err != nil {
			panic(err)
		}

		findOne := options.FindOne()
		town, err := countriesCollection.FindOne(bson.D{{"_id", objID}}, findOne)
		if err != nil {
			panic(err)
		}
		fmt.Println(town)

		// find One with object ID and return Name of the city only
		findOne.SetProjection(bson.D{{"name", 1}})
		town, err = countriesCollection.FindOne(bson.D{{"_id", objID}}, findOne)
		if err != nil {
			panic(err)
		}

		// find One With population in range
		// $lt = less than
		// $gt = greater than
		findOne.SetProjection(bson.D{{"name", 1}, {"population", 1}})
		town, err = countriesCollection.FindOne(bson.M{
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
		town, err = countriesCollection.FindOne(bson.M{"famousfor": regex}, findOne)
		if err != nil {
			panic(err)
		}
		fmt.Println(town)
		fmt.Println()
		fmt.Println("###################################################################################")
	*/
}

func InsertTown(countriesCollection *country.CountriesCollection) (ids []string, err error) {

	if isDocumentAlreadyExist(countriesCollection) {
		return
	}

	countries := []country.Country{}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ids, err = countriesCollection.InsertMany(countries)
	}()
	wg.Wait()
	if err != nil {
		return
	}

	return
}

func isDocumentAlreadyExist(countriesCollection *country.CountriesCollection) bool {
	towns, find_err := countriesCollection.Find(options.Find())
	if find_err != nil {
		panic(find_err)
	}
	if towns != nil {
		return true
	}
	return false
}
