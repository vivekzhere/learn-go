package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type Item struct {
	Title	string
	URL	string
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	resp, err := http.Get("http://reddit.com/r/golang.json")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		log.Fatal(err)
	}

	for _, child := range r.Data.Children {
		fmt.Println(child.Data.Title)
	}
}

