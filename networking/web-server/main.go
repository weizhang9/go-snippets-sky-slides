package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Jeremiah was a bullfrog"))
	})
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err.Error())
	}
}
