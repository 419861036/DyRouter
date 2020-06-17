package core

import (
	"fmt"
	"io"
	"net/http"
)

type Dispath interface {
	Handler(r *http.Request, w http.ResponseWriter, fn func(*http.Request, http.ResponseWriter))
}

type Proxy struct {
}

func (p Proxy) Handler(r *http.Request, w http.ResponseWriter, fn func(*http.Request, http.ResponseWriter)) {
	fmt.Printf("Received request %s %s %s %s\n", r.Method, r.Host, r.RemoteAddr, r.URL.String())
	transport := http.DefaultTransport
	res, err := transport.RoundTrip(r)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	if res.StatusCode == http.StatusForbidden || res.StatusCode == http.StatusInternalServerError {
		HttpProxy{}.Handler(r, w, fn)
		return
	}

	// step 3
	for key, value := range res.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	fn(r, w)
	w.WriteHeader(res.StatusCode)

	io.Copy(w, res.Body)
	res.Body.Close()
	/*proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ModifyResponse = func(response *http.Response) error {
		response.Header.Add("good", "11")
		return nil
	}
	proxy.ServeHTTP(w, r)*/
}

type HttpProxy struct {
}

func (p HttpProxy) Handler(r *http.Request, w http.ResponseWriter, fn func(*http.Request, http.ResponseWriter)) {
	client := &http.Client{}
	req, _ := http.NewRequest(r.Method, r.URL.Scheme+"://"+r.URL.Host+r.URL.Path, nil)
	resp, _ := client.Do(req)
	//回调
	fn(r, w)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

type StreamProxy struct {
}

func (s StreamProxy) Handler(r *http.Request, w http.ResponseWriter, fn func(*http.Request, http.ResponseWriter)) {

}
