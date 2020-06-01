package http

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	wg.Done()
}

func init() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.0:8000", nil)
}
func Handler(wg sync.WaitGroup) {
	wg = wg
	wg.Add(1)
}
