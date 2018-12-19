package main

import (
	"fmt"
)

const (
	// each constant below will be restored as a bit in a total of 8 bit (because there are 8 constants)
	isUKTeam          = 1 << iota // 0000 0001
	isEuropeanTeam                // 0000 0010
	isInPremierLeague             // 0000 0100
	isPlayingTonight              // 0000 1000
	isLiveOnSky                   // 0001 0000
	isWon                         // 0010 0000
	isPromoting                   // 0100 0000
	hasTransferNews               // 1000 0000
)

func main() {
	// 0000 0001 |    0000 0100      |    0000 1000     |  0001 0000  = 0001 1101
	var Tottenham byte = isUKTeam | isInPremierLeague | isPlayingTonight | isLiveOnSky
	// now we have stored all these data of a team in a single byte (8 bits)
	fmt.Printf("%b\n", Tottenham)

	// if constant isUKTeam is set on variable Tottenham, isUKTeam & Tottenham will give the value of isUKTeam
	// then we compare isUKTeam & Tottenham with isUKTeam, if it's set, the result will be true
	fmt.Printf("Is Tottenham a UK Team? %v\n", isUKTeam&Tottenham == isUKTeam)
	fmt.Printf("Is Tottenham a European Team? %v\n", isEuropeanTeam&Tottenham == isEuropeanTeam)
	fmt.Printf("Is Tottenham in Premier League? %v\n", isInPremierLeague&Tottenham == isInPremierLeague)
	fmt.Printf("Is Tottenham playing tonight? %v\n", isPlayingTonight&Tottenham == isPlayingTonight)
	fmt.Printf("Is Tottenham match Live on Sky? %v\n", isLiveOnSky&Tottenham == isLiveOnSky)
	fmt.Printf("Did Tottenham win the match today (probably not)? %v\n", isWon&Tottenham == isWon)
	fmt.Printf("Are we promoting Tottenham (probably not)? %v\n", isPromoting&Tottenham == isPromoting)
	fmt.Printf("Does Tottenham has transfer news today? %v\n", hasTransferNews&Tottenham == hasTransferNews)
}
