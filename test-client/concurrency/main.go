package main

import (
	"context"
	"fmt"
	"time"

	lib "github.com/garcios/asset-trak-portfolio/lib/concurrency"
)

func main() {
	ctx := context.Background()
	g, ctx := lib.WithContext(ctx, 3) // Limit to 3 concurrent goroutines

	for i := 0; i < 10; i++ {
		g.Go(func() error {
			fmt.Println("Starting goroutine")
			time.Sleep(1 * time.Second)
			fmt.Println("Finishing goroutine")
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("All goroutines finished")

}
