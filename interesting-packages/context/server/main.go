// oh yeah, context also works for http communication 🤠
// cancellation can be sent across servers
// but values can't from http package, can implement with own code tho
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"go-snippets-sky-slides/interesting-packages/context/log"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8086", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "Handler started")
	defer log.Println(ctx, "Handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hola!")
	case <-ctx.Done():
		log.Println(ctx, ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
