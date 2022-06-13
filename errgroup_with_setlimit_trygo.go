// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("Hello, 世界")
	var eg errgroup.Group
	eg.SetLimit(2)
	eg.TryGo(func() error {
		for {
			fmt.Printf("From go 1\n")
			time.Sleep(1 * time.Second)
		}
	})

	eg.TryGo(func() error {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("From go 2\n")
		}
		return nil
	})

	eg.TryGo(func() error {
		for {
			fmt.Printf("From go 3\n")
		}
	})

	_ = eg.Wait()
}
