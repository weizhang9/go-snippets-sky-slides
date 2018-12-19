package main

import "fmt"

func main() {
	// last in, first out
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	// when defer a function, it takes the arguments at the calling time, not executing time
	a := 1
	defer fmt.Println(a) // a is 1
	a = 2
}
