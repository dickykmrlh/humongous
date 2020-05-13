package country

import (
	"fmt"
	"sync"

	"github.com/dickymrlh/humongous/domain/country"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PlayAroundWithCountry(countriesCollection *country.CountriesCollection) {
	fmt.Println("###################################################################################")
	err := InsertTown(countriesCollection)
	if err != nil {
		panic(err)
	}

	// find With elemMatch
	// It specifies that if a document (or nested document)
	// matches all of our criteria
	countries, err := countriesCollection.Find(bson.M{
		"exports.foods": bson.M{
			"$elemMatch": bson.M{
				"name":  "bacon",
				"tasty": true,
			},
		},
	}, options.Find())
	if err != nil {
		panic(err)
	}
	fmt.Println(countries)
	fmt.Println()
	fmt.Println("###################################################################################")

	// find With or
	countries, err = countriesCollection.Find(bson.M{
		"$or": []bson.M{
			bson.M{"_id": "mx"},
			bson.M{"name": "United States"},
		},
	}, options.Find().SetProjection(bson.D{{"_id", 1}}))
	if err != nil {
		panic(err)
	}
	fmt.Println(countries)
	fmt.Println()
	fmt.Println("###################################################################################")

	// remove with bad bacon
	removedCount, err := countriesCollection.Remove(bson.M{
		"exports.foods": bson.M{
			"$elemMatch": bson.M{
				"name":  "bacon",
				"tasty": false,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(removedCount)
	fmt.Println()
	fmt.Println("###################################################################################")
}

func InsertTown(countriesCollection *country.CountriesCollection) error {

	if isDocumentAlreadyExist(countriesCollection) {
		return nil
	}

	countries := []country.Country{
		country.Country{
			ID:   "us",
			Name: "United States",
			Exports: country.Export{
				[]country.Food{
					country.Food{Name: "bacon", Tasty: true},
					country.Food{Name: "burgers"},
				},
			},
		},
		country.Country{
			ID:   "ca",
			Name: "Canada",
			Exports: country.Export{
				[]country.Food{
					country.Food{Name: "bacon", Tasty: false},
					country.Food{Name: "syrup", Tasty: true},
				},
			},
		},
		country.Country{
			ID:   "mx",
			Name: "Mexico",
			Exports: country.Export{
				[]country.Food{
					country.Food{Name: "salsa", Tasty: true, Condiment: true},
				},
			},
		},
	}

	var wg sync.WaitGroup
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = countriesCollection.InsertMany(countries)
	}()
	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}

func isDocumentAlreadyExist(countriesCollection *country.CountriesCollection) bool {
	towns, find_err := countriesCollection.Find(bson.D{}, options.Find())
	if find_err != nil {
		panic(find_err)
	}
	if towns != nil {
		return true
	}
	return false
}
