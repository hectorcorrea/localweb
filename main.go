package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var webAddress = ""
var path = "public"

func main() {
	port := flag.Int("port", 9001, "Listening port")
	flag.Parse()

	// Make sure the localFolder exists
	localFolder := "./" + path
	if err := initDir(localFolder); err != nil {
		log.Printf("Error: %s", err)
	}

	// Setup the request handlers
	urlPath := "/" + path + "/"
	fs := http.FileServer(http.Dir(localFolder))
	http.Handle(urlPath, http.StripPrefix(urlPath, fs))
	http.HandleFunc("/", home)

	// Start the web server
	webAddress = fmt.Sprintf("localhost:%d", *port)
	log.Printf("Listening for requests at: http://%s", webAddress)
	log.Printf("Serving files from: %s", localFolder)
	err := http.ListenAndServe(webAddress, nil)
	if err != nil {
		log.Fatal("Failed to start the web server: ", err)
	}
}

func home(resp http.ResponseWriter, req *http.Request) {
	url := "http://" + webAddress + "/" + path + "/"
	html := "<html><body>"
	html += "<p>Sample web server for static files.</p>"
	html += "<p>Serves anything under "
	html += "<a href=" + url + ">" + url + "</a></p>"
	html += "</body></html>"
	fmt.Fprint(resp, html)
}

func dirExists(path string) bool {
	// source: https://stackoverflow.com/a/10510783/446681
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func initDir(path string) error {
	if dirExists(path) {
		return nil
	}

	// Create the folder...
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		return err
	}

	// and put a sample file on it
	filename := path + "/sample.html"
	content := []byte("<h1>Hello World</h1>")
	err := ioutil.WriteFile(filename, content, 0644)
	return err
}
