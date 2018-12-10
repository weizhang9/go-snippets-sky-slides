package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://news.sky.com")
	if err != nil {
		log.Fatal(err.Error)
	}
	defer res.Body.Close()
	sn, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error)
	}
	fmt.Printf("%s", sn)
}