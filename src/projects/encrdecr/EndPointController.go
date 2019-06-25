package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	setupEndpointHandlers()
	startServer()
}

func setupEndpointHandlers() {
	http.HandleFunc("/info", infoHandler)
}

func infoHandler(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, "This is encryption decryption application")
}

func startServer() {
	port := ":2607"
	http.ListenAndServe(port, nil)
	fmt.Printf("\nStarted server on port %s\n", port)
}
