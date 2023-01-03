package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./pkg/http/web/app/dist"))
	http.Handle("/", fs)
	log.Println("Listing for requests with Vue UI at http://localhost:8000/")

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
