package main

import (
	"fmt"
	"strconv"
)

func main() {
	t := sum("42", 9, 17, 7, 24, 8, 10)
	fmt.Println(t)

	ta, err := sumAddr("42", 9, 17, 7, 24, 8, 10)
	if err != nil {
		fmt.Println("You failed:", err)
	}
	fmt.Println(ta)
}

// since I named my return value "total", variable total will be returned if nothing else is specified
// named returns are initialized to the zero values for their types when the function begins
func sum(baseNum string, e ...int) (total int) {
	total, _ = strconv.Atoi(baseNum)
	for _, v := range e {
		total += v
	}
	return
}

// return value type can be a pointer
// when returning a pointer, the memory will be promoted to the heap automatically
// you can return as many values as you like in Go functions
func sumAddr(baseNum string, e ...int) (totalAddr *int, err error) {
	b, err := strconv.Atoi(baseNum)
	if err != nil {
		return nil, fmt.Errorf("Cannot convert %v to integer", baseNum)
	}
	for _, v := range e {
		b += v
		totalAddr = &b
	}
	return
}
