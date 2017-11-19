package main

import (
	"fmt"
	"sync"
)

// Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawled is the map for all crawled data
type Crawled struct {
	urls  map[string]int
	mutex sync.Mutex
}

var crawled Crawled

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	crawled.mutex.Lock()
	_, found := crawled.urls[url]
	crawled.urls[url]++
	crawled.mutex.Unlock()
	if found {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	var waitGroup sync.WaitGroup
	for _, u := range urls {
		waitGroup.Add(1)
		go func(url string) {
			defer waitGroup.Done()
			Crawl(url, depth-1, fetcher)
		}(u)
	}
	waitGroup.Wait()
	return
}

func main() {
	crawled = Crawled{urls: make(map[string]int)}
	Crawl("http://golang.org/", 4, fetcher)

	for url, occurances := range crawled.urls {
		fmt.Println("url=", url, "occurances=", occurances)
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
