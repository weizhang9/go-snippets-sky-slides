package main

import (
	"fmt"
)

func main() {
	f2 := Add2()
	fmt.Printf("Calling Add2() for 3 gives %v\n", f2(3))

	f := Add(2)
	fmt.Printf("Calling Add(2) for 3 gives %v", f(3))

	// when self-invoking closures within a loop
	// best practice is to pass outter variables to closures to avoid strange behaviours during async operations
	for i := 0; i < 10; i++ {
		func(i int) {
			fmt.Println(i)
		}(i) // now changes on the outter scope variable i won't affect this inner scope operation
	}

}

// closures/lambda/function literal/anonymous functions can be self-invoked or assigned to a variable
// or passed in as a parameter, or as a return value
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Add(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
