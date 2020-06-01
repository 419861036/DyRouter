package main

import (
	"./com/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	http.Handler(wg)
	wg.Wait()
}
