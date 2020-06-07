package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mallvielfrass/HDesktopBox/filmix"
)

type msg string
type Profile struct {
	Name    string
	Hobbies []string
}

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../front/index.html")
}
func search(w http.ResponseWriter, r *http.Request) {
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	token := filmix.GetToken()
	fmt.Println("token:", token)
	api := filmix.API(token)
	filmId := api.Search("lorax")
	fmt.Println(filmId.Items[0].ID)
	IdInf := api.Info(filmId.Items[0].ID)
	fmt.Printf("%s | %s | %d | %d \n", IdInf.OriginalTitle, IdInf.Title, IdInf.Year, IdInf.ID)
	js, err := json.Marshal(IdInf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func main() {
	//	msgHandler := msg("Hello from Web Server in Go")
	fmt.Println("Server is listening...")
	http.HandleFunc("/s", search)
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8181", nil)
}
