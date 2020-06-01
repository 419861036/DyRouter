package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendData(r *http.Request) []byte {
	client := &http.Client{}
	fmt.Println(r.URL.Path)
	req, err := http.NewRequest(r.Method, "http://www.baidu.com", strings.NewReader(""))
	if err != nil {
		// handle error
	}
	fmt.Println(r.Method)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return body
}
