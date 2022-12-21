/* Using the default mux */

package main

import (
	"encoding/json"
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

type Middleware func(http.HandlerFunc) http.HandlerFunc

func chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

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

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products = []Product{
	{Id: 101, Name: "Pen", Cost: 10},
	{Id: 102, Name: "Pencil", Cost: 20},
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(products); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var newProduct Product
		if err := decoder.Decode(&newProduct); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		newProduct.Id = len(products) + 101
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(newProduct); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		fmt.Fprintf(w, "Products handler executed")
	}
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintln(w, "Customer requests are processed")
}

func main() {
	/*
		http.HandleFunc("/", profile(log(indexHandler)))
		http.HandleFunc("/products", profile(log(productsHandler)))
		http.HandleFunc("/customers", profile(log(customersHandler)))
	*/
	http.HandleFunc("/", chain(indexHandler, log, profile))
	http.HandleFunc("/products", chain(productsHandler, log, profile))
	http.HandleFunc("/customers", chain(customersHandler, log, profile))
	http.ListenAndServe(":8080", nil)
}
