package main

import "fmt"

func main() {
	i := intCounter(-9)
	var inc Incrementer = &i
	for i := 0; i < 17; i++ {
		fmt.Println(inc.increment())
	}
}

type Incrementer interface {
	increment() int
}

type intCounter int

// interfaces are implemented implicitly in Go by attach the method(s) in contract to a type
func (ic *intCounter) increment() int {
	*ic++
	return int(*ic)
}
