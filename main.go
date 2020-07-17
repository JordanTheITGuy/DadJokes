package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Joke struct {
	ID string `json:"id"`
	JOKE string `json:"joke"`
	STATUS int `json:"status"`
}

func main() {
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	joke1 := Joke{}
	jsonErr := json.Unmarshal(bytes,&joke1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(joke1.JOKE)
}