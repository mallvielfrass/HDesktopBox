package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

type Countries struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type Directors struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type FilmInfo struct {
	IDKinopoisk   int64       `json:"idKinopoisk"`
	Slogan        string      `json:"slogan"`
	Title         string      `json:"title"`
	Countries     interface{} `json:"countries"`
	Poster        string      `json:"poster"`
	Directors     interface{} `json:"directors"`
	ShortStory    string      `json:"short_story"`
	Year          int64       `json:"year"`
	OriginalTitle string      `json:"original_title"`
}
type Result struct {
	FilmInfo []FilmInfo `json:"filminfo"`
	Count    int        `json:"count"`
}

func search(w http.ResponseWriter, r *http.Request) {
	//profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	//token := filmix.GetToken()
	//fmt.Println("token:", token)
	keys := r.URL.Query()
	//deviceGUID := keys.Get("deviceGUID") //Get return empty string if key not found
	q := keys.Get("q")
	lim, err := strconv.Atoi(keys.Get("l"))
	if err != nil {
		log.Println("lim err")
	}
	if len(keys.Get("l")) < 1 {
		log.Println("Url Param 'limit' is missing")
		lim = 10

	}
	fmt.Println("lim=", lim)
	fmt.Println("q=", q)
	api := filmix.API("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6ImZ4LTVlZGQ3MGNiMzUzZGEifQ.eyJpc3MiOiJodHRwczpcL1wvZmlsbWl4Lm1lIiwiYXVkIjoiaHR0cHM6XC9cL2ZpbG1peC5tZSIsImp0aSI6ImZ4LTVlZGQ3MGNiMzUzZGEiLCJpYXQiOjE1OTE1NzA2MzUsIm5iZiI6MTU5MTU1OTgzNSwiZXhwIjoxNTk0MTYyNjM1LCJwYXJ0bmVyX2lkIjoiMiIsImhhc2giOiI2NzVhZjZiMDBiYWZhMDhmOGYwMDE5Y2Q3YWMyYmM4Zjk0MmQ0NDY5IiwidXNlcl9pZCI6bnVsbCwiaXNfcHJvIjpmYWxzZSwiaXNfcHJvX3BsdXMiOmZhbHNlLCJzZXJ2ZXIiOiIifQ.0xnppIMMr53upxHhrNbPkD0QJ5I14EyG72qMxfnbQL4")
	filmId := api.Search("lorax")
	count := len(filmId.Items)
	//	a := []FilmInfo{}
	// equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println("api ", filmId)
	//IdInf := api.Info(filmId.Items[0].ID)
	//
	//fmt.Printf("%s | %s | %d | %d \n", IdInf.OriginalTitle, IdInf.Title, IdInf.Year, IdInf.ID)
	if count < lim {
		lim = count
	}
	fmt.Printf("count %d \n", lim)
	FilmInfostr := []FilmInfo{}
	for i := 0; i < lim; i++ {
		IdInf := api.Info(filmId.Items[i].ID)
		fmt.Printf("%s | %s | %d | %d \n", IdInf.OriginalTitle, IdInf.Title, IdInf.Year, IdInf.ID)
		b := FilmInfo{
			IDKinopoisk:   IdInf.IDKinopoisk,
			Slogan:        IdInf.Slogan,
			Title:         IdInf.Title,
			Countries:     IdInf.Countries,
			Poster:        IdInf.Poster,
			Directors:     IdInf.Directors,
			ShortStory:    IdInf.ShortStory,
			Year:          IdInf.Year,
			OriginalTitle: IdInf.OriginalTitle,
		}
		FilmInfostr = append(FilmInfostr, b)
	}
	res := Result{
		FilmInfo: FilmInfostr,
		Count:    count,
	}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func main() {
	//	msgHandler := msg("Hello from Web Server in Go")

	http.HandleFunc("/s", search)
	http.HandleFunc("/", index)
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
