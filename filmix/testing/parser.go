package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Series представляет сериал
type Series struct {
	Seasons     []Season `json:"seasons"` // Список сезонов
	Translation string   `json:"name"`    // Вариант локализации
}

// Seasons представляет сезон сериала
type Season struct {
	Num      int                `json:"season"`   // Порядковый номер сезона
	Episodes map[string]Episode `json:"episodes"` // Эпизоды сезона (map[код_эпизода)Описание_эпизода)
}

// Seasons представляет эпизод сезона
type Episode struct {
	Num         int    `json:"episode"` // Порядковый номер эпизода
	AdSkip      int    `json:"ad_skip"`
	Title       string `json:"title"`    // Название эпизода
	ReleasedStr string `json:"released"` // Дата релиза в формате ГГГГ-ММ-ДД
	Files       []File `json:"files"`    // Список файлов
}

// File представляет файл эпизода
type File struct {
	Quality int    `json:"quality"` // Качество
	ProPlus bool   `json:"proPlus"`
	URL     string `json:"url"` // URL к файлу
}

// --

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("dat2.json")
	check(err)

	var allSeries []Series
	err = json.Unmarshal(dat, &allSeries)
	check(err)

	for _, series := range allSeries {
		fmt.Printf("Локализация от %s\n", series.Translation)
		for _, season := range series.Seasons { // тип map[string]Episode не сохраняет порядок, себе нужно будет это обработать самому
			for episodeCode, episode := range season.Episodes {
				fmt.Printf("Сезон %d, эпизод %s: %s\n", season.Num, episodeCode, episode.Title)
			}
		}

		fmt.Println("")
	}
}
