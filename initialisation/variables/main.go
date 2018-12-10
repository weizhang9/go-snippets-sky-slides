package main

import (
	"fmt"
)

// declare (and optionally assign) variables in a block
// export a variable by capitalising first letter
// exported variables are accessible from other packages, i.e. public
// declared global variables won't throw error if not been used
// but declared local variables have to be used otherwise compiler will error
var (
	SquadName         = "Cosmonauts"
	SqaudMotto        = "Challenging trumps boring"
	SquadPlayerNumber = 11
	SquadError        = false
	squadXmasParty    complex128
)

func main() {
	// declare a var, then assign a value
	// useful for a variable needing a larger scope than where its value can be assigned
	var s string
	s = "tomato"
	fmt.Println(s)
	// reassign a new value, but cannot re-declared in the same scope
	s = "potato"
	fmt.Println(s)

	// declare a var and assign a value in one go
	// useful for controlling the variable's type
	var i float64 = 42
	fmt.Println(i)

	// use shorthand assignment operator (:=) to declare and assign
	// let the compiler infer the type
	// := cannot be used outside of func
	b := true
	fmt.Printf("%v, %T", b, b)
}
