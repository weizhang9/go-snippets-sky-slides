package main

import (
	"fmt"
	"reflect" // https://golang.org/pkg/reflect/ useful for looking up variable's type
)

type I interface {
	M()
}

type T1 struct{ watevs string }

func (T1) M() {}

type T2 struct{}

func (T2) M() {}

func main() {
	var i I = T1{"not empty"}
	fmt.Println("i belongs to package:", reflect.TypeOf(i).PkgPath())
	fmt.Println("name of i's type is:", reflect.TypeOf(i).Name())
	fmt.Printf("string representation of i's type is: %v\n\n", reflect.TypeOf(i).String())

	i = T2{}
	fmt.Println("i belongs to package:", reflect.TypeOf(i).PkgPath())
	fmt.Println("name of i's type is:", reflect.TypeOf(i).Name())
	fmt.Println("string representation of i's type is:", reflect.TypeOf(i).String())
}
