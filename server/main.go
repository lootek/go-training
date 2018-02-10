package main

import (
	"fmt"
	"net/http"
)

func greetFn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[greetFn] Welcome to %v!", r.URL)
}

type myGreet struct{}

func (myGreet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[myGreet.ServeHTTP] Welcome to %v!", r.URL)
}

func main() {
	http.HandleFunc("/foo", greetFn)

	greet := myGreet{}
	http.Handle("/bar", greet)

	// https://golang.org/pkg/net/http
	fmt.Println("Staring server...")
	http.ListenAndServe(":80", nil)
}
