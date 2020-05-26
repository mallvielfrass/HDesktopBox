package main

import (
	"fmt"

	"github.com/mallvielfrass/HDesktopBox"
)

func main() {
	token := HDesktopBox.GetToken()
	api := HDesktopBox.API(token)
	filmId := api.Search("lorax")
	fmt.Println(filmId.Items[0].ID)
	IdInf := api.Info(filmId.Items[0].ID)
	fmt.Printf("%s | %s | %d | %d \n", IdInf.OriginalTitle, IdInf.Title, IdInf.Year, IdInf.ID)
}
