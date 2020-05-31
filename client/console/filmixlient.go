package main

import (
	"fmt"

	"github.com/mallvielfrass/HDesktopBox/filmix"
)

func main() {
	token := filmix.GetToken()
	api := filmix.API(token)
	filmId := api.Search("lorax")
	fmt.Println(filmId.Items[0].ID)
	IdInf := api.Info(filmId.Items[0].ID)
	fmt.Printf("%s | %s | %d | %d \n", IdInf.OriginalTitle, IdInf.Title, IdInf.Year, IdInf.ID)

}
