package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 9001, "Listening port")
	flag.Parse()

	filePath := "."
	urlPath := "/"

	// Setup the request handler
	fs := http.FileServer(http.Dir(filePath))
	http.Handle(urlPath, http.StripPrefix(urlPath, fs))

	// Start the web server
	webAddress := fmt.Sprintf("localhost:%d", *port)
	log.Printf("Listening for requests at: http://%s", webAddress)
	log.Printf("Serving files from: %s", filePath)
	err := http.ListenAndServe(webAddress, nil)
	if err != nil {
		log.Fatal("Failed to start the web server: ", err)
	}
}
