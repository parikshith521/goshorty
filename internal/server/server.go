package server

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the organized Go server! 🚀")
}

func Run(addr string) {
	http.HandleFunc("/", handler)
	log.Printf("Server starting on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
