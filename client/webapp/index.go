package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index/index.html")
	//fmt.Fprint(w, "index")
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/test")
	http.ServeFile(w, r, "./static/test/index.html")
	//fmt.Fprint(w, "index")
}
