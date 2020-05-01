package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dickymrlh/humongous/domain/town"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c := town.GetInstance(ctx)

	t := []town.Town{
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

	id, err := c.Insert(ctx, t)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
