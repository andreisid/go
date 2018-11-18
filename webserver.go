package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range files {
		fmt.Fprint(w, i.Name(), "\n")
	}
}

func main1() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
