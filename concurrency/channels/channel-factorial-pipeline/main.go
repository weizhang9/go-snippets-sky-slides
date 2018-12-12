// sorry for not adding notes on concurrency code snippets; 
// this is a more advanced concept that would require researching and learning to obsorb, 
// trying to wrote notes on it would end up into an essay, 
// please feel free to chat with me if you have any questions about my code
package main

import (
	"fmt"
)

func main() {
	numC := num()
	f := factorial(numC)
	for n := range f {
		fmt.Println(n)
	}
}

func num() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for k := 6; k < 16; k++ {
				c <- k
			}
		}
		close(c)
	}()
	return c
}

func factorial(c <-chan int) <-chan int {
	d := make(chan int)
	go func() {
		for n := range c {
			d <- multiply(n)
		}
		close(d)
	}()
	return d
}

func multiply(n int) int {
	total := 1
	for l := n; l > 0; l-- {
		total *= l
	}
	return total
}
