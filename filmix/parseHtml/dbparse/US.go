package main

import (
	// import third party libraries

	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

func getFilmFromPage(url string, filmBox MainBox) MainBox {
	res, err := http.Get("https://filmix.co" + url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	doc.Find("#dle-content .film a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		//	linkText := linkTag.Text()
		href, _ := linkTag.Attr("href")
		s := strings.Split(href, "/")
		id := strings.Split(s[len(s)-1], "-")[0]
		Name := linkTag.Find(".film-name").Text()
		sr := strings.NewReader(Name)
		tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
		buf, err := ioutil.ReadAll(tr)
		if err != err {
			// обработка ошибки
		}
		Names := string(buf)
		//	fmt.Printf("Film #%d: id: %s - %s, link %s \n", index, id, Names, href)
		filmBox.AddItem(Item{
			ID:   id,
			Name: Names,
			Link: href,
		})
	})
	return filmBox
}
func (box *MainBox) AddItem(item Item) []Item {
	box.Item = append(box.Item, item)
	return box.Item
}

type MainBox struct {
	Item []Item
}
type Item struct {
	ID   string
	Name string
	Link string
}

func main() {
	fmt.Println("привет")
	ItemList := []Item{}
	boxClear := MainBox{ItemList}
	box := getFilmFromPage("/catalog/%DE", boxClear)

	Lbox := len(box.Item)
	db, err := sql.Open("sqlite3", "symbolTables.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	dbEx := ""
	for i := 0; i < Lbox; i++ {
		bID, err := strconv.Atoi(box.Item[i].ID)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		str := fmt.Sprintf("insert or replace into T28 (id, Name, Url) values (%d, '%s','%s');\n",
			bID, box.Item[i].Name, box.Item[i].Link)
		dbEx = dbEx + str
	}
	fmt.Println(dbEx)
	_, err = db.Exec(dbEx)
	if err != nil {
		panic(err)
	}
}
