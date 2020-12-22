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

	counter := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Received request from " + r.RemoteAddr)

		counter++
		if counter >= 5 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "You've hit %v, %v\n", hostname, counter)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
