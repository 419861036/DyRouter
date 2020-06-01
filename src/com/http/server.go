package http

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "hello world")
	// filter request header
	body := SendData(r)
	//filter repsonse header

	fmt.Fprintln(w, string(body))

}

func init() {
	fmt.Println("开启代理服务器：8000")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)

}
func Handler(wg *sync.WaitGroup) {
	wg1 = *wg
	for {
		time.Sleep(5 * time.Second)
	}
}
