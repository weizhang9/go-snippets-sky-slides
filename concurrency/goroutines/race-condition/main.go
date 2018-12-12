// sorry for not adding notes on concurrency code snippets; 
// this is a more advanced concept that would require researching and learning to obsorb, 
// trying to wrote notes on it would end up into an essay, 
// please feel free to chat with me if you have any questions regarding my code
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Miyoko:")
	go incrementor("Reiko:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "\tCounter:", counter)
	}
	wg.Done()
}

// go run -race main.go to check if race condition is existed in the code
