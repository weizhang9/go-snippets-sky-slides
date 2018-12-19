package main

import "fmt"

type I interface {
	M()
}

type T struct{} // T implements interface I

func (T) M() {}

func main() {
	var t *T // t is declared but not assigned any value, so it's nil
	if t == nil {
		fmt.Println("t is nil")
	} else {
		fmt.Println("t is not nil")
	}

	var i I = t // even i is assigned to an empty struct's pointer, interface type variable's value is consisted of its value and its type, since it has a dynamic type of struct, it's not nil
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is not nil")
	}

	// 👆🏼 this is important as your if condition could behave exactly the opposite as you intended
}
