package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dickymrlh/humongous/domain/town"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c := town.GetInstance(ctx)

	t := []town.Town{
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
	var id []string
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		id, err = c.Insert(ctx, t)
	}()
	wg.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
