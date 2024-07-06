package main

import (
	"fmt"
	"net/http"
)

func rootHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, you've requested from %s\n", r.URL)
}

func main() {
	http.HandleFunc("/", rootHanlder)
	http.ListenAndServe(":3000", nil)
}
