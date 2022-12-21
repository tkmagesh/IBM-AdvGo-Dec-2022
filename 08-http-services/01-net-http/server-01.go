package main

import (
	"fmt"
	"net/http"
)

type server struct {
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s\n", r.URL, r.Method)
	/*
		fmt.Fprintln(w, "Hello World!")
	*/
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/customers":
		fmt.Fprintln(w, "Customer requests are processed")
	case "/products":
		fmt.Fprintln(w, "Product requests are processed")
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	s := &server{}
	http.ListenAndServe(":8080", s)
}
