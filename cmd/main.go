package main

import (
	"context"
	"time"

	"github.com/dickymrlh/humongous/domain/town"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	town.GetInstance(ctx)
}
