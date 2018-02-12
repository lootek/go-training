package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lootek/go-training/server/storage"
)

type myGreet struct{}

func (myGreet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[myGreet.ServeHTTP] Welcome to %v!", r.URL)
}

func datastoreHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		item := struct {
			Key   string `json:"k"`
			Value int    `json:"v"`
		}{}

		json.NewDecoder(r.Body).Decode(&item)
		fmt.Printf("%#v\n", item)

		storage.Save(item.Key, item.Value)
	case http.MethodGet:
		var values []int

		if keys, ok := r.URL.Query()["key"]; ok {
			values = storage.Get(keys...)
		} else {
			values = storage.Get()
		}

		fmt.Printf("%#v\n", values)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		json.NewEncoder(w).Encode(&values)
	}
}

func setupStorage() {
	storage.Run()
}

func setupServer() {
	greet := myGreet{}
	http.Handle("/foo", greet)

	http.HandleFunc("/data", datastoreHandle)

	// https://golang.org/pkg/net/http
	fmt.Println("Staring server...")
	http.ListenAndServe(":80", nil)
}

func main() {
	setupStorage()
	setupServer()
}
