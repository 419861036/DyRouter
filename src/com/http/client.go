package http

import (
	"../config"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func SendData(r *http.Request, w http.ResponseWriter) {
	fmt.Printf("Received request %s %s %s %s\n", r.Method, r.Host, r.RemoteAddr, r.URL.String())
	servers := config.GLOBAL_CONFIG.Servers[0]
	remote, err := url.Parse(servers.Location)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ModifyResponse = func(response *http.Response) error {
		response.Header.Add("good", "11")
		return nil
	}
	proxy.ServeHTTP(w, r)

}
