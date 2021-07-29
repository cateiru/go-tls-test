package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://www.googleapis.com/oauth2/v1/certs"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
