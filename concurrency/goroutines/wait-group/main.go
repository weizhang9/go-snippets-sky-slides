package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(-1))
	wg.Add(2)
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
