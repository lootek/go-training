package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lootek/go-training/cli/json"
	"github.com/lootek/go-training/cli/xml"
	"github.com/pkg/errors"
)

const (
	host = "http://127.0.0.1"
	port = 80
)

func readData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		err = errors.Wrap(err, "cannot make a request")
		log.Println(err)
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

type ArrayUnmarshaler interface {
	Unmarshal([]byte) ([]int, error)
}

func write(u ArrayUnmarshaler, data []byte) {
	fmt.Println(u.Unmarshal(data))
}

func main() {
	fmt.Println("Hello")

	urls := []string{
		fmt.Sprintf("%s:%v/data", host, port),
		fmt.Sprintf("%s:%v/data?format=xml", host, port),
	}

	for _, u := range urls {
		fmt.Printf("\n%s\n", u)

		data, err := readData(u)
		if err != nil {
			log.Println(err)
			return
		}

		xu := xml.XMLUnmarshaler{}
		write(xu, data)

		ju := &json.JSONUnmarshaler{}
		write(ju, data)
	}
}
