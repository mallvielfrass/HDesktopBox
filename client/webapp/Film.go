package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mallvielfrass/HDesktopBox/filmix"
)

func film(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	//token := filmix.GetToken()
	api := filmix.API("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6ImZ4LTVlZGQ3MGNiMzUzZGEifQ.eyJpc3MiOiJodHRwczpcL1wvZmlsbWl4Lm1lIiwiYXVkIjoiaHR0cHM6XC9cL2ZpbG1peC5tZSIsImp0aSI6ImZ4LTVlZGQ3MGNiMzUzZGEiLCJpYXQiOjE1OTE1NzA2MzUsIm5iZiI6MTU5MTU1OTgzNSwiZXhwIjoxNTk0MTYyNjM1LCJwYXJ0bmVyX2lkIjoiMiIsImhhc2giOiI2NzVhZjZiMDBiYWZhMDhmOGYwMDE5Y2Q3YWMyYmM4Zjk0MmQ0NDY5IiwidXNlcl9pZCI6bnVsbCwiaXNfcHJvIjpmYWxzZSwiaXNfcHJvX3BsdXMiOmZhbHNlLCJzZXJ2ZXIiOiIifQ.0xnppIMMr53upxHhrNbPkD0QJ5I14EyG72qMxfnbQL4")
	//filmId := api.Search("lorax")
	//fmt.Println(filmId.Items[0].ID)
	IdInf := api.Info(i)
	fmt.Printf("%s | %s | %d | %d \n", IdInf.OriginalTitle, IdInf.Title, IdInf.Year, IdInf.ID)
	fmt.Fprint(w, IdInf)
}
