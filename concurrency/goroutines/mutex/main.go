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

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Miyoko:")
	go incrementor("Reiko:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "\tCounter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
