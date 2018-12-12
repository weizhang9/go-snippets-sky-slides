// sorry for not adding notes on concurrency code snippets; 
// this is a more advanced concept that would require researching and learning to obsorb, 
// trying to wrote notes on it would end up into an essay, 
// please feel free to chat with me if you have any questions regarding my code
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Miyoko:")
	go incrementor("Reiko:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter))
	}
	wg.Done()
}
