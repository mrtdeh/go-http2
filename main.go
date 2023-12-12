package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP/2!")
	})

	server := &http.Server{
		Addr: ":8080",
	}

	http2.ConfigureServer(server, &http2.Server{})

	log.Fatal(server.ListenAndServeTLS("certs/server.crt", "certs/server.key"))
}
