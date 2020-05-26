package kinopub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestToken struct {
	Code            string `json:"code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	Interval        int    `json:"interval"`
	ExpiresIn       int    `json:"expires_in"`
}

func Register() RequestToken {

	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://api.service-kp.com/oauth2/device?grant_type=device_code&client_id=xbmc&client_secret=cgg3gtifu46urtfp2zp1nqtba0k2ezxh", nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.Status)
	//fmt.Println(string(body))
	var data RequestToken
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	return data
}
