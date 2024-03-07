package main

import (
	"fmt"
	"net/http"
)

func counter() func(w http.ResponseWriter, r *http.Request) {
	var views int
	return func(w http.ResponseWriter, r *http.Request) {
		views++
		fmt.Fprintf(w, "<h1>views: %d<h1>", views)

		fmt.Println("views: ", views)
	}
}

func main() {
	addr := "localhost:2024"
	fmt.Printf("http://%s \n", addr)

	http.HandleFunc("/index", counter())

	http.ListenAndServe(addr, nil)
}
