package main

import "fmt"

type I interface {
	M()
}
type T1 struct{ watevs string }

func (T1) M() {}

type T2 struct{}

func (T2) M() {}

func main() {
	var i I = T1{"not empty"} // i is an interface type, i have a static type (interface I) and a dynamic type (type of value assigned to i, T1 in this instance)
	fmt.Printf("value of i: %v\t type of i: %T\n", i, i)
	i = T2{} // since T2 also implement interface I, we can reassign i to T2 and cast type/value to T2
	fmt.Printf("value of i: %v\t\t type of i: %T", i, i)
}
