package main

import "fmt"

func main() {
	// allow to use initialisation for tag
	switch balance := 15 + 17; balance {
	// allow to have multiple tests in a single case
	case 10, 25, 32:
		fmt.Println("You could afford Sky Cinema (after you bought Sky TV), Sky Fibre Unlimited or Sky Mobile")
		// switch statement defaults to break after a condition is met
		// to un-break, use keyword fallthrough, but this keyword does NOT care about logic
		fallthrough
	case 40:
		fmt.Println("You could afford Sky Broadband & TV")
		// break
		fmt.Println("If you need to leave a case early, use break above it")
	case 45:
		fmt.Println("You could afford Sky TV Deal")
	default:
		fmt.Println("You could give your money to Sky in some other ways")
	}
}
