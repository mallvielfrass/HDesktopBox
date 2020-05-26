package filmix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Directors struct
type Directors struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Actors struct
type Actors struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Countries struct
type Countries struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//Genres struct
type Genres struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//Info struct
type Info struct {
	Actors          []Actors    `json:"actors"`
	Category        string      `json:"category"`
	Countries       []Countries `json:"countries"`
	Directors       []Directors `json:"directors"`
	Duration        int64       `json:"duration"`
	Genres          []Genres    `json:"genres"`
	ID              int64       `json:"id"`
	IDKinopoisk     int64       `json:"idKinopoisk"`
	Mpaa            string      `json:"mpaa"`
	OriginalTitle   string      `json:"original_title"`
	Poster          string      `json:"poster"`
	Quality         string      `json:"quality"`
	RatingImdb      float64     `json:"ratingImdb"`
	RatingKinopoisk float64     `json:"ratingKinopoisk"`
	ShortStory      string      `json:"short_story"`
	Slogan          string      `json:"slogan"`
	Title           string      `json:"title"`
	Updated         string      `json:"updated"`
	URL             string      `json:"url"`
	VotesImdb       int64       `json:"votesImdb"`
	VotesKinopoisk  int64       `json:"votesKinopoisk"`
	VotesNeg        int64       `json:"votesNeg"`
	VotesPos        int64       `json:"votesPos"`
	Year            int64       `json:"year"`
}

//Info return info about film
func (api *StructAPI) Info(id int64) Info {

	idString := strconv.FormatInt(id, 10)
	compl := api.Host + "/film/" + idString + "/details"
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
	//fmt.Println(resp.Status)
	//fmt.Println(string(body))
	var infData Info
	if err := json.Unmarshal(body, &infData); err != nil {
		panic(err)
	}
	return infData
}
