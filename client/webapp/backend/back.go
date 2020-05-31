package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../front/index.html")
}
func main() {
	//	msgHandler := msg("Hello from Web Server in Go")
	fmt.Println("Server is listening...")
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8181", nil)
}
