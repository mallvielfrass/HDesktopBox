package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mallvielfrass/HDesktopBox/filmix"
)

type msg string
type Profile struct {
	Name    string
	Hobbies []string
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
	filmId := api.Search(q)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	//	fmt.Println(js)
	w.Write(js)
}
func jearch(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	//	fmt.Println(js)
	fmt.Fprintf(w, `{"filminfo":[{"idKinopoisk":470036,"slogan":"«From the creators of Despicable Me»","title":"Лоракс","countries":[{"id":2,"name":"США"}],"poster":"https://thumbs.bookdline.live/s/posters/thumbs/w220/loraks_film_2012_37834_0.jpeg","directors":[{"id":"Kristoffer-Boae","name":"Кристоффер Боэ"}],"short_story":"Так, много лет спустя, по вине человека, с планеты исчезли все деревья. Воздух теперь можно покупать в отдельных бутылочках, а внешний мир радует своей красотой. Одно желание - и у тебя под окном океан, второй - а там лес. Одно отличие от прошлой жизни, что все это пластиковое. Но многие дети, не видели настоящих деревьев, и потому уже не знают, чем этот мир может быть хуже. Теду всего двенадцать лет, но у него уже есть девушка. Одри, которую он так сильно любит, любит мечтать, а еще ее самое большое желание - увидеть живое дерево. Самое настоящее, которое видели еще ее родители. У Теда нет желания сказать, Одри нет, потому сначала, он спрашивает у бабушки, а потом отправляется к знакомому, который видел последнее дерево. И как оказывается, именно по его вине все это произошло, когда-то давно мужчине не послушал духа леса Лоракса и уничтожил все деревья. Теперь Теду предстоит посадить последнее дерево и найти духа, но это будет сложно, так как против этого выступают и власти и родители мальчика.","year":2012,"original_title":"Dr. Seuss\\' The Lorax"},{"idKinopoisk":341898,"slogan":"-","title":"Лоракс","countries":[{"id":2,"name":"США"}],"poster":"https://thumbs.bookdline.live/s/posters/thumbs/w220/loraks-the-lorax-1972_77557_0.jpeg","directors":[{"id":"Xouli-Praett","name":"Хоули Прэтт"}],"short_story":"Мальчик попал  в безлесную пустошь, чтобы встретиться с разорившимся промышленником, который  рассказал ему историю о произошедшем с ним. Сначала он создал процветающий бизнес, основанный на бесполезном продукте моды, который был получен из лесных деревьев.","year":1972,"original_title":"The Lorax"}],"count":2}`)
}

type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
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
func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./front/index.html")
	//fmt.Fprint(w, "index")

}
func main() {
	//	msgHandler := msg("Hello from Web Server in Go")
	fileServer := http.FileServer(FileSystem{http.Dir("static")})
	http.Handle("/static/", http.StripPrefix(strings.TrimRight("/static/", "/"), fileServer))
	http.HandleFunc("/s", search)
	http.HandleFunc("/j", jearch)
	http.HandleFunc("/", index)
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
