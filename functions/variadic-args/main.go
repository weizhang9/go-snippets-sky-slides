package main

import (
	"fmt"
	"strconv"
)

func main() {
	// A variadic parameter (indicate with ...) accepts zero or more values of a specified type
	// Go compiler will convert these values into a slice of the specified type
	sum("42", 9, 17, 7, 24, 8, 10)
}

// each func can only have one variadic param, and the variadic param needs to be the last param otherwise runtime will smack you
func sum(baseNum string, e ...int) {
	b, _ := strconv.Atoi(baseNum) // use strconv.Atoi to convert string to integer
	for _, v := range e {         // throw away index of passed-in integer(s)
		b += v
	}
	fmt.Println(b)
}
