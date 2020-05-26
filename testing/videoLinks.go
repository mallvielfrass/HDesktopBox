package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func link(id int64) []byte {

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6ImZ4LTVlYjliNDhlMTAwMWUifQ.eyJpc3MiOiJodHRwczpcL1wvZmlsbWl4Lm1lIiwiYXVkIjoiaHR0cHM6XC9cL2ZpbG1peC5tZSIsImp0aSI6ImZ4LTVlYjliNDhlMTAwMWUiLCJpYXQiOjE1ODkyMjg2ODYsIm5iZiI6MTU4OTIxNzg4NiwiZXhwIjoxNTkxODIwNjg2LCJwYXJ0bmVyX2lkIjoiMiIsImhhc2giOiI2NzVhZjZiMDBiYWZhMDhmOGYwMDE5Y2Q3YWMyYmM4Zjk0MmQ0NDY5IiwidXNlcl9pZCI6bnVsbCwiaXNfcHJvIjpmYWxzZSwiaXNfcHJvX3BsdXMiOmZhbHNlLCJzZXJ2ZXIiOiIifQ.KumFD0InHANhf3LrO7FLRTjCvl3xzT0n7xbruIEnSTg"
	serv := "http://5.61.48.15/partner_api/video_links/"

	//par := "search=" + word + "&sort=news_read&page=0"
	s := strconv.FormatInt(id, 10)
	compl := serv + s //+ "/details"

	client := &http.Client{}

	req, _ := http.NewRequest("GET", compl, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-FX-Token", token)
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	return body
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Create("dat2")
	check(err)
	defer f.Close()
	f.Write([]byte(link(7412)))
	f.Sync()
}
