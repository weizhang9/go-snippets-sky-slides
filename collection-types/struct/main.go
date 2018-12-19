package main

import "fmt"

// declare a struct
// struct's fields don't have to be same type
// but non-blank field names must be unique
type randomInfo struct {
	Movies              // embedding a struct
	upTime, downTime    int
	FavGames, FavAnimes []string
}

type Movies struct {
	Name map[string]bool
}

func main() {
	L := randomInfo{
		// there is no inheritance when embedding a struct
		Movies:    Movies{map[string]bool{"Harry Potter": false, "Star Wars": false, "Interstellar": true}},
		upTime:    86400,
		downTime:  1440,
		FavGames:  []string{"Battlefield V", "RDR2", "Stardew Valley"},
		FavAnimes: []string{"银魂", "‎死亡筆記", "夏目友人帳", "攻殻機動隊"},
	}

	// struct is value type, value won't mutated unless passing a pointer
	copy := L
	// can assign value with dot syntax as well
	copy.Name = map[string]bool{"Terminator": true, "Star Wars": false, "Interstellar": true}

	fmt.Println(L.Movies)
	fmt.Println(copy.Movies)

	anonymousStruct := struct{ name string }{name: "WALL-E"}
	fmt.Printf("\nI'm %v, an anonymous struct; come in handy when you want to organise data in a struct but not want to create a formal type.", anonymousStruct.name)
}
