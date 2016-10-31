package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
)

type Item struct {
	Title	string
	URL	string
}
func (i Item) String() string {
	return fmt.Sprintf("%s\n%s", i.Title, i.URL)
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	items, err := Get("golang")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}

func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(r.Data.Children))

	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}
