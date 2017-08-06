package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello secure World on this day 6th August 2017!")
}

func main() {
	http.HandleFunc("/echo", hello)
	http.ListenAndServe(":8081", nil)
}
