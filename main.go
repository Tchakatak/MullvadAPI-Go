package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type data struct {
	IP            string `json:"ip"`
	Country       string `json:"country"`
	MullvadExitIP bool   `json:"mullvad_exit_ip"`
}

func main() {

	url := "https://am.i.mullvad.net/json"

	mullvadClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Go-Api-Req")

	res, getErr := mullvadClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	data1 := data{}
	jsonErr := json.Unmarshal(body, &data1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if data1.MullvadExitIP == true {

		fmt.Println("Your currently connected to mullvad")
	} else {
		fmt.Println("Your currently not connected to mullvad")
	}
	fmt.Println("Your Curent IP address is : ", data1.IP)
	fmt.Println("Your Current exit server is in :", data1.Country)
}
