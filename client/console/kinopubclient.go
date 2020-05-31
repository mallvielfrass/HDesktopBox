package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/mallvielfrass/HDesktopBox/kinopub"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func oscheck(err error) {
	if err != nil {
		fmt.Printf("err %s", err)
		os.Exit(1)
	}
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type InitStruct struct {
	Code            string `json:"code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	Interval        int    `json:"interval"`
	ExpiresIn       int    `json:"expires_in"`
}
type FilmixTomlConfig struct {
	Code            string `toml:"Code"`
	UserCode        string `toml:"UserCode"`
	VerificationURI string `toml:"VerificationURI"`
	Interval        int    `toml:"Interval"`
	ExpiresIn       int    `toml:"ExpiresIn"`
}

func initFilmix() {
	regData := kinopub.Register()
	fmt.Println("creating config kinopub.toml")
	f, err := os.Create("kinopub.toml")
	check(err)
	defer f.Close()
	fstring := fmt.Sprintf("Code = \"%s\"\nUserCode = \"%s\"\nVerificationURI = \"%s\"\nInterval = %d\nExpiresIn = %d\n ",
		regData.Code, regData.UserCode, regData.VerificationURI, regData.Interval, regData.ExpiresIn)
	_, err = f.WriteString(fstring)
	check(err)
	//fmt.Println(regData.Code)

}
func main() {
	if fileExists("kinopub.toml") {
		var config FilmixTomlConfig
		if _, err := toml.DecodeFile("kinopub.toml", &config); err != nil {
			fmt.Println(err)

		}

	} else {
		initFilmix()

	}
}
