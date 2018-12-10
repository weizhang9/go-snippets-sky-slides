package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case p := <-ch:
		fmt.Println("Job done", p)

	case <-ctx.Done():
		fmt.Println("Moving on")
	}
}
