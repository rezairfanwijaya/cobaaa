package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.ServeMux{}
	mux.Handle("/", home)

	if err := http.ListenAndServe(":8080", &mux); err != nil {
		log.Fatal(err)
	}
}

var home http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hallo")
}
