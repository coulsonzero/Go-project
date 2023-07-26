package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go study(ctx)

	time.Sleep(time.Second * 5)
}

func study(ctx context.Context) error {
	i := 1
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Printf("study %d seconds. \n", i)
		}
		i++
	}
}

/*
study 1 seconds.
study 2 seconds.
study 3 seconds.
study 4 seconds.
*/
