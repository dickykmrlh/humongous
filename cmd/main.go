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
			Name:       "Oakland",
			Population: 390724,
			LastCensus: time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local),
			FamousFor:  []string{"Fried chicken", "Mahershala Ali"},
			Mayor:      town.Politican{Name: "Libby Schaaf "},
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
