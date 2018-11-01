package server

import (
	"fmt"
	"log"
	"net/http"
)

func handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the url was %v", r.URL.Path)
}

func Run() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
