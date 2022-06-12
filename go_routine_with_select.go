// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	channel := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	var sg sync.WaitGroup
	sg.Add(1)
	go func() {
		defer sg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("####### context done go 1######\n")
				return
			case msg := <-channel:
				fmt.Printf("####### %v ######\n", msg)

			}

		}

	}()

	sg.Add(1)
	go func() {
		defer sg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("####### context done go 2######\n")
				return
			case channel <- rand.Intn(1000):
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	sg.Wait()
}
