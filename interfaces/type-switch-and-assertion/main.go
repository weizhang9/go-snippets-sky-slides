package main

import "fmt"

type I1 interface {
	M()
}

type I2 interface {
	I1
	N()
}

type T struct {
	name string
}

func (T) M() {}

func (T) N() {}

func main() {
	// use type assertion to convert the concrete type of an interface type to another interface type
	// the concrete must satisfy the target interface type
	var v1 I1 = T{"woof"}
	var v2 I2
	fmt.Printf("type of v2 is %T\n", v2)
	v2, ok := v1.(I2) // syntax: baseInterfaceType.(targetInterfaceType)
	if !ok {
		fmt.Printf("Type assertion failed\n\n")
	} else {
		fmt.Printf("type of v2 is %T\n\n", v2)
	}

	// use switch to check variable's type
	// we can do this to interface type as they have a static type (interface) and a dynamic type
	// the dynamic type is determined by the type of the assigned value to the variable
	// which mean any type that implements the interface
	var s interface{} = "I will determine the dynamic type of the empty interface"
	switch s.(type) {
	case int:
		fmt.Println("The concrete type is an int")
	case bool:
		fmt.Println("The concrete type is a boolean")
	case string:
		fmt.Println("The concrete type is string")
	default:
		fmt.Println("The concrete type is unknown")
	}
}
