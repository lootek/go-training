package storage

import "fmt"

type writeRequest struct {
	key   string
	value int
}

type readRequest struct {
	keys     []string
	response chan []int
}

var (
	readReq  chan readRequest
	writeReq chan writeRequest
)

func Run() {
	readReq = make(chan readRequest)
	writeReq = make(chan writeRequest)

	go func() {
		datastore := map[string]int{}

		for {
			select {
			case r := <-readReq:
				fmt.Printf("%#v\n", r)

				values := make([]int, 0)

				if len(r.keys) > 0 {
					for _, k := range r.keys {
						values = append(values, datastore[k])
					}
				} else {
					for _, v := range datastore {
						values = append(values, v)
					}
				}

				r.response <- values
			case w := <-writeReq:
				fmt.Printf("%#v\n", w)

				datastore[w.key] = w.value
			}
		}
	}()
}

func Get(keys ...string) []int {
	resp := make(chan []int)

	readReq <- readRequest{keys, resp}

	values := <-resp
	return values
}

func Save(key string, val int) {
	writeReq <- writeRequest{key, val}
}
