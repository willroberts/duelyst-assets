package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://play.duelyst.com/api/me/inventory/card_collection"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))
}
