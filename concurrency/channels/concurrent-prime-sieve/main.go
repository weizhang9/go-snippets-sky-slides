// A concurrent prime sieve
// this is actually an example from Rob Pike
package main

import "fmt"

// daisy-chain Filter processes
func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)

		out := make(chan int)
		go Filter(ch, out, prime)
		ch = out
	}
}

// send the sequence 2, 3, 4, ... to channel 'ch'
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// copy the values from channel 'in' to channel 'out'
// removing those divisible by 'prime'
func Filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
