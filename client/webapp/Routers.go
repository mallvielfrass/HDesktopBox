package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func faviconRouter(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/logo/favicon.ico")
}
func CssRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	object := vars["object"]
	//response := fmt.Sprintf("%s", id)
	path := "./static/" + page + "/css/" + object
	fmt.Println(path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, path)
}
func JsRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	object := vars["object"]
	//response := fmt.Sprintf("%s", id)
	path := "./static/" + page + "/js/" + object
	fmt.Println(path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, path)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Choose the folder to serve
	staticDir := "/static/"
	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	return router
}
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
