package main

import (
	"fmt"
	"math"
)

// methods can be defined for any named type except a pointer or an interface (not to confused with pointer receiver)
func main() {
	// a method on an embedded struct can be called directly on an instance of the embedding type
	// if the embedding type has a method of the same name, it will override the one on embedded type
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.getHypotenuse())
}

type Point struct {
	x, y float64
}

// define a method on pointer receiver of type Pointer who is an embedded type in NamedPoint below
func (p *Point) getHypotenuse() float64 {	
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

type NamedPoint struct {	// the embedding type in this example
	Point
	name string
}