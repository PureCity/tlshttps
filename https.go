package main

import (
	"fmt"
	"net/http"
)

const (
	SERVER_PORT       = 8080
	SERVER_DEMAIN     = "localhost"
	RESPONSE_TEMPLATE = "hello"
)

func main() {
	server := fmt.Sprintf("%s:%d/", SERVER_DEMAIN, SERVER_PORT)
	http.HandleFunc(server, rootHandler)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
}

// Show root handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}
