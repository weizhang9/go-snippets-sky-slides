// sorry for not adding notes on concurrency code snippets;
// this is a more advanced concept that would require researching and learning to obsorb,
// trying to wrote notes on it would end up into an essay,
// please feel free to chat with me if you have any questions regarding my code
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {
	// it's common to increase the default thread which is the same number as your processors
	// but you would need to do some testing to find the breaking point
	// so that your app runs faster
	// while not overloading go scheduler to deal with more goroutines than it can handle
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// you can query your CPU number with runtime package
	fmt.Println(runtime.NumCPU())
	// runtime.GOMAXPROCS(int) lets you set the number of threads you would like to use for your app
	// when setting it as -1, you are asking it to return the previous set value
	fmt.Println(runtime.GOMAXPROCS(-1))
	wg.Add(2)
	// i'm only launching 2 goroutines here
	// but it's very common to have tens of thousands of goroutines in a loop to make it fly
	go reiko()
	go miyoko()
	wg.Wait()
}

func miyoko() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Miyoko:\t", i)
	}
	wg.Done()
}

func reiko() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Reiko:\t", i)
	}
	wg.Done()
}
