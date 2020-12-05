package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()

	log.Printf("Server started on ... %s", hostname)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Received request from " + r.RemoteAddr)

		fmt.Fprintf(w, "You've hit %v\n", hostname)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
