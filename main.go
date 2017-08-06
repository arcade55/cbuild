package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello secure World!")
}

func main() {
	http.HandleFunc("/echo", hello)
	http.ListenAndServe(":8081", nil)
}
