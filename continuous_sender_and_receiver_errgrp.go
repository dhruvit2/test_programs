// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

// 2 go routine
// 1 chan
// random number

// run for 1 sec and terminate
func main() {

	comChan := make(chan int)
	defer close(comChan)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Millisecond)
	eg, ctx := errgroup.WithContext(ctx)

	// producer
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case comChan <- rand.Intn(1000):
				time.Sleep(1 * time.Millisecond)
			}
		}
		return nil
	})

	// consumer
	eg.Go(func() error {
		for {
			select {
			case recNum := <-comChan:
				fmt.Printf("[%v] Received random number %v \n", time.Now(), recNum)
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		return nil
	})

	_ = eg.Wait()
}
