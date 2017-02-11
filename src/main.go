package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Pet struct {
	Name    string
	Hobbies []string
	Likes   int
}

func main() {
	resp, err := http.Get("http://localhost:9000/list")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pets := []Pet{}
	error := json.Unmarshal(data, &pets)

	if error != nil {
		log.Fatal(err)
	}

	for _, pet := range pets {
		fmt.Println(pet)
	}

}
