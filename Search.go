package HDesktopBox

//par := "search=" + word + "&sort=news_read&page=0"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Items struct {
	Actors          []Actors    `json:"actors"`
	Category        string      `json:"category"`
	Countries       []Countries `json:"countries"`
	Genres          []Genres    `json:"genres"`
	ID              int64       `json:"id"`
	OriginalTitle   string      `json:"original_title"`
	Poster          string      `json:"poster"`
	Quality         string      `json:"quality"`
	RatingImdb      float64     `json:"ratingImdb"`
	RatingKinopoisk float64     `json:"ratingKinopoisk"`
	Title           string      `json:"title"`
	Updated         string      `json:"updated"`
	VotesNeg        int64       `json:"votesNeg"`
	VotesPos        int64       `json:"votesPos"`
	Year            int64       `json:"year"`
}
type Req struct {
	HasNextPage bool    `json:"has_next_page"`
	Items       []Items `json:"items"`
	Page        int64   `json:"page"`
	Status      string  `json:"status"`
}

//Search info on filmix
func (api *StructAPI) Search(word string) Req {
	serv := "http://5.61.48.15/partner_api/list?"
	par := "search=" + word + "&sort=news_read&page=0"
	compl := serv + par
	client := &http.Client{}

	req, _ := http.NewRequest("GET", compl, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-FX-Token", api.FXToken)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data Req
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	return data
}
