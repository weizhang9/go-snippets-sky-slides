This repo is corresponding to the `Introduction to Go` slides

To run code in this repo locally, first install go from https://golang.org/doc/install on your machine and set up necessary env vars, then clone this repo to your machine.

Each directory that contains a `main.go` should be treated as an individual package (of course they are not reusable packages, they are just code snippets as an illustrations to fundamental go code charastics), therefore you could `cd` into any package and run command `go run main.go` or specifiy the path to the `main.go` from your current location after `go run` (e.g. if you are in root directory of the repo and would like to run `data-store-in-bits/main.go`, run `go run awesome-bit-operations/data-store-in-bits/main.go` or `go build awesome-bit-operations/data-store-in-bits/main.go` to build the binary first and execute the binary instead)

If you don't want to run it locally, you can click into the go playground links on the slides, and run it there; thought playground can't show multi-processor goroutines, or any server operations.