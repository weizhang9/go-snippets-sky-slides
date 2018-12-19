// https://golang.org/pkg/context/
// communicate (propagation) cancellation and context value
// saving resources and making requests more efficient
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	// time package allows to sleep while giving a receive channel to use !!!
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func main() {
	// init a ctx (context.Background)
	ctx := context.Background()
	// init a new ctx with cancel feature (context.WithCancel)
	ctx, cancel := context.WithCancel(ctx)
	// now if we type anything in terminal, the request is cancelled
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	// can use time.AfterFunc for arranged cancellation
	// now cancel will be called after 2 second if no stdin was grabbed before that
	time.AfterFunc(2*time.Second, cancel)

	// init a new ctx with timeout to achieve above (context.WithTimeOut)
	// but we will need to call the cancel regardless to clear the allocated memory
	newCtx, newCancel := context.WithTimeout(ctx, 10*time.Second)
	defer newCancel()

	sleepAndTalk(ctx, 5*time.Second, "hello")
	sleepAndTalk(newCtx, 5*time.Second, "goodbye")
}
