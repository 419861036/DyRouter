package http

import (
	"../config"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// filter request header
	SendData(r, w)
	//filter repsonse header

	//fmt.Fprintln(w, string(body))

}

func init() {
	port := strconv.Itoa(config.GLOBAL_CONFIG.Port)
	fmt.Println("开启代理服务器：" + port)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:"+port, nil)

}
func Handler(wg *sync.WaitGroup) {
	wg1 = *wg
	for {
		time.Sleep(5 * time.Second)
	}
}
