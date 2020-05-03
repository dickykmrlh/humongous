package main

import (
	"context"
	"time"

	"github.com/dickymrlh/humongous/domain/town"
	t "github.com/dickymrlh/humongous/usecase/town"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	townCollection := town.GetInstance(ctx)
	t.PlayAroundWithTown(townCollection)
}
