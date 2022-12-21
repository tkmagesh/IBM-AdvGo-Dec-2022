/* Using the default mux */

package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
func log(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.URL, r.Method)
		handler(w, r)
	}
}
*/

//middleware
func log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.URL, r.Method)
		handler(w, r)
	}
}

func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		fmt.Println("Time taken : ", elapsed)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprintln(w, "Product requests are processed")
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintln(w, "Customer requests are processed")
}

func main() {
	http.HandleFunc("/", profile(log(indexHandler)))
	http.HandleFunc("/products", profile(log(productsHandler)))
	http.HandleFunc("/customers", profile(log(customersHandler)))
	http.ListenAndServe(":8080", nil)
}
