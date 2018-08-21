package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type safeCache struct {
	m   map[string]bool
	mux sync.Mutex
}

var cache safeCache

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, exit chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	// for i := 0; i < 4-depth; i++ {
	// 	fmt.Printf("\t")
	// }
	// fmt.Printf("Start: %s\n", url)

	defer func() {
		exit <- url
	}()
	if depth <= 0 {
		return
	}

	cache.mux.Lock()
	if _, ok := cache.m[url]; ok {
		cache.mux.Unlock()
		return
	}
	cache.m[url] = true
	cache.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	// _, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	childExits := make(chan string, len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, childExits)
	}
	for range urls {
		// txt:= <-childExits
		// for i := 0; i < 5-depth; i++ {
		// 	fmt.Printf("\t")
		// }
		// fmt.Printf("Exit: %s\n", txt)
		<-childExits
	}
	return
}

func main() {
	cache = safeCache{m: make(map[string]bool)}
	exit := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, exit)
	// fmt.Printf("Exit: %s\n", <-exit)
	<-exit
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
