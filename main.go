package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Received request from " + r.RemoteAddr)

		name, _ := os.Hostname()
		fmt.Fprintf(w, "You've hit %v\n", name)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
