package rest

import (
	"fmt"
	"net/http"
	"sync"
)

func StartHTTPServer(listenPort int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Printf("Starting HTTP server on port %d\n...", listenPort)
	http.HandleFunc("/", Get_Index_handler)
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}
