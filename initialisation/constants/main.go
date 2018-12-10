package main

import (
	"fmt"
)

// constants can only be numbers, characters (runes), strings or booleans
// naming convention same as variables (PascalCase for exported, camelCase for local)
const mcD = "lorem ipsum latte"

func main() {
	// strings are made of a collection of bytes/runes
	// but we need to explicitly cast the value into byte type
	bs := []byte(mcD)
	fmt.Println(bs, len(bs) == len(mcD))
	// UTF8 can be used anywhere in the code, not just string
	薇 := []byte("薇")
	fmt.Println(薇)
	fmt.Println([]byte("🤪"))

	// constants can be untyped
	// their type will then be inferred by compiler, depending on the assigned value
	// which makes it possible for implicit conversion during evalution with variables
	// if the underlying types are compatible
	const i = 42
	fmt.Printf("%T\n", i)
	s := 66.6
	fmt.Printf("Value: %v\t Type: %T\n", i+s, i+s)
	// variables have to be typed, so need to be explicitly converted for evalution
	v := 9
	fmt.Println(float64(v) + s)

	//constants are "immutable", but we can shadow the value
	fmt.Printf("McDougal's is no longer serving %v, ", mcD)
	// constant mcD now has value of "jägerbombs"
	// this is because the inner scope's value takes priority
	mcD := "jägerbombs"
	fmt.Printf("We are now serving %v", mcD)
}
