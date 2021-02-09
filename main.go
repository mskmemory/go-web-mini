package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.ListenAndServe(":8080", r)
}
