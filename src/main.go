package main

import (
	"./com/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go http.Handler(&wg)
	wg.Wait()
}
