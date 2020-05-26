package filmix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//StructAPI retun FXToken - auth hdbox(http://5.61.48.15/partner_api) token
type StructAPI struct {
	Host    string
	FXToken string
}

//API entry point
func API(FXToken string) *StructAPI {
	return &StructAPI{
		Host:    "http://5.61.48.15/partner_api",
		FXToken: FXToken,
	}
}

//GetTokenStruct struct for GetToken
type GetTokenStruct struct {
	Token   string `json:"token"`
	Expired int    `json:"expired"`
}

//GetToken get FXToken
func GetToken() (FXToken string) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://filmix.vielfrass.tk", nil)
	//	req, _ := http.NewRequest("POST", url, nil)
	//req.Header.Add("Accept", "application/x-www-form-urlencoded")
	//	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.Status)
	//fmt.Println(string(body))
	var data GetTokenStruct
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	FXToken = data.Token
	return FXToken
}
