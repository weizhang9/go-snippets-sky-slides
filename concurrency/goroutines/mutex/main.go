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
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
