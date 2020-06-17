package http

import (
	"../config"
	"../core"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// filter request header

	core.Handdler(w, r)
}

func init() {
	http.HandleFunc("/", IndexHandler)
	for _, server := range config.GLOBAL_CONFIG.Servers {
		go func(server config.Server) {
			p := strconv.Itoa(server.Port)
			log.Println("开启代理服务器成功,端口：" + p)

			http.ListenAndServe(":"+p, nil)
		}(server)

	}

}
func Handler(wg *sync.WaitGroup) {
	wg1 = *wg
	for {
		time.Sleep(5 * time.Second)
	}
}
