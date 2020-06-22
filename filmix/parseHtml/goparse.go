// file: list_posts.go
package main

import (
	// import third party libraries
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getPage(url string) string {
	doc, err := goquery.NewDocument("https://filmix.co" + url)
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
	fmt.Println("Pages:")
	for i := 0; i < (len(catalog)); i++ {
		//fmt.Println("i:", i)
		fmt.Printf("i: %d, pages: %s\n", i, getPage(catalog[i]))
		time.Sleep(3000 * time.Millisecond)
	}
}

func main() {
	postScrape()
}
