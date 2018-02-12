package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
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

		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			log.Printf("cannot decode json: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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

		switch r.URL.Query().Get("format") {
		case "xml":
			data, err := xml.Marshal(&values)
			if err != nil {
				log.Printf("cannot encode json: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/xml; charset=utf-8")
			w.Write(data)
		case "json":
			fallthrough
		default:
			data, err := json.Marshal(&values)
			if err != nil {
				log.Printf("cannot encode json: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write(data)
		}
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
