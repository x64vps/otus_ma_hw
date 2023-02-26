package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)

	_ = http.ListenAndServe(":8000", mux)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, "{'status': 'OK'}")
}
