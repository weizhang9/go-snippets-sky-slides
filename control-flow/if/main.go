package main

import "fmt"

func main() {
	skyProducts := map[string]int{
		"Sky TV Deal":         45,
		"Sky Cinema":          10,
		"Sky Broadband & TV":  40,
		"Sky Fibre Unlimited": 25,
		"Sky Mobile":          32,
	}
	
	// if initialisation; condition in one line
	if p, ok := skyProducts["Sky Fibre Unlimited"]; ok {
		fmt.Printf("Sky Fibre Unlimited is £%v per month for 18 months", p)
	}
}
