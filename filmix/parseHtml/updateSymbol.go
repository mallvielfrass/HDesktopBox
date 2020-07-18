package main

import (
	// import third party libraries
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

func getPage(url string) string {
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

	max := ""
	doc.Find(".navigation a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		//link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		max = linkText
		//fmt.Printf("Pages #%d: '%s' - '%s'\n", index, linkText, link)
	})
	if max == "" {
		max = "1"
	}
	return max
}
func GetFilm() {
	fmt.Println("Films:")
	f, e := os.Open("index.html")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	doc, e := goquery.NewDocumentFromReader(f)
	if e != nil {
		log.Fatal(e)
	}
	doc.Find("#dle-content .film a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		//	linkText := linkTag.Text()
		href, _ := linkTag.Attr("href")
		s := strings.Split(href, "/")
		id := strings.Split(s[len(s)-1], "-")[0]
		link := linkTag.Find(".film-name").Text()

		fmt.Printf("Film #%d: id: %s - '%s'\n", index, id, link)
	})
}
func postScrape() {

	catalog := map[int]string{
		0:  "/catalog/0-9", //0-9
		1:  "/catalog/%C0", //А
		2:  "/catalog/%C1", //Б
		3:  "/catalog/%C2", //В
		4:  "/catalog/%C3", //Г
		5:  "/catalog/%C4", //Д
		6:  "/catalog/%C5", //Е
		7:  "/catalog/%C6", //Ж
		8:  "/catalog/%C7", //З
		9:  "/catalog/%C8", //И
		10: "/catalog/%C9", //Й
		11: "/catalog/%CA", //К
		12: "/catalog/%CB", //Л
		13: "/catalog/%CC", //М
		14: "/catalog/%CD", //Н
		15: "/catalog/%CE", //О
		16: "/catalog/%CF", //П
		17: "/catalog/%D0", //Р
		18: "/catalog/%D1", //С
		19: "/catalog/%D2", //Т
		20: "/catalog/%D3", //У
		21: "/catalog/%D4", //Ф
		22: "/catalog/%D5", //Х
		23: "/catalog/%D6", //Ц
		24: "/catalog/%D7", //Ч
		25: "/catalog/%D8", //Ш
		26: "/catalog/%D9", //Щ
		27: "/catalog/%DD", //Э
		28: "/catalog/%DE", //Ю
		29: "/catalog/%DF", //Я
	}
	catalogHuman := map[int]string{
		0:  "/catalog/0-9", //0-9
		1:  "/catalog/А",   //А
		2:  "/catalog/Б",   //Б
		3:  "/catalog/В",   //В
		4:  "/catalog/Г",   //Г
		5:  "/catalog/Д",   //Д
		6:  "/catalog/Е",   //Е
		7:  "/catalog/Ж",   //Ж
		8:  "/catalog/З",   //З
		9:  "/catalog/И",   //И
		10: "/catalog/Й",   //Й
		11: "/catalog/К",   //К
		12: "/catalog/Л",   //Л
		13: "/catalog/М",   //М
		14: "/catalog/Н",   //Н
		15: "/catalog/О",   //О
		16: "/catalog/П",   //П
		17: "/catalog/Р",   //Р
		18: "/catalog/С",   //С
		19: "/catalog/Т",   //Т
		20: "/catalog//У",  //У
		21: "/catalog/Ф",   //Ф
		22: "/catalog/Х",   //Х
		23: "/catalog/Ц",   //Ц
		24: "/catalog/Ч",   //Ч
		25: "/catalog/Ш",   //Ш
		26: "/catalog/Щ",   //Щ
		27: "/catalog/Э",   //Э
		28: "/catalog/Ю",   //Ю
		29: "/catalog/Я",   //Я
	}
	db, err := sql.Open("sqlite3", "symb.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Pages:")

	for i := 0; i < (len(catalog)); i++ {
		//fmt.Println("i:", i)
		dat := getPage(catalog[i])
		fmt.Printf("i: %s, pages: %s, link: %s\n", catalogHuman[i], dat, catalog[i])
		pages, err := strconv.Atoi(dat)
		if err == nil {
			fmt.Println(pages)
		}
		_, err = db.Exec("insert or replace into alphabet (id, name, valueFilm,url,pages) values ($1, $2,'0',$3,$4)",
			i, catalogHuman[i], catalog[i], pages)
		if err != nil {
			panic(err)
		}
		//fmt.Println(result.LastInsertId())
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	postScrape()
}
