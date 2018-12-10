package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	needy bool
}

type cat struct {
	animal
	aloof bool
}

// every type implements empty interface
// and empty interface is also a type that can be assigned to a variable or as self-defined type
type I interface{} 

func specs(a I) {	// accept any type of arguments
	fmt.Println(a)
}

func main() {
	doggo := dog{animal{"woof"}, true}
	catto := cat{animal{"meow"}, true}
	specs(doggo)
	specs(catto)
	
	// define a slice of empty interface type
	dc := []interface{}{doggo, catto, "just popping in", 42.0}
	specs(dc)	
}
