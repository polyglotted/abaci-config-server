package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/data")))

	log.Printf("Serving /data on HTTP port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
