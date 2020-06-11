package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//file routers
	router.HandleFunc("/{page}/css/{object}", CssRouter)
	router.HandleFunc("/{page}/js/{object}", JsRouter)
	router.HandleFunc("/favicon.ico", faviconRouter)

	//api
	router.HandleFunc("/s", search)
	router.HandleFunc("/j", jearch)

	//page
	router.HandleFunc("/film/{id}", film)
	router.HandleFunc("/test", test)
	router.HandleFunc("/", index)

	//run serve
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
