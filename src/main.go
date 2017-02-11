package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Pet struct {
	Name    string
	Hobbies []string
	Likes   int
}

type Pets []Pet

func (p Pets) Len() int           { return len(p) }
func (p Pets) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pets) Less(i, j int) bool { return p[i].Name > p[j].Name }

func main() {
	resp, err := http.Get("http://localhost:9000/list")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pets := Pets{}
	err = json.Unmarshal(data, &pets)

	sort.Sort(pets)

	if err != nil {
		log.Fatal(err)
	}

	for _, pet := range pets {
		fmt.Println(pet)
	}

}
