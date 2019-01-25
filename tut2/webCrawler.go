package main

import (
	"fmt"
  "sync"
)

type SafeMap struct {
  m map[string]bool
  mux sync.Mutex
}

// returns whether or not key is in safe map
func (sm *SafeMap) Found(key string) bool {
  sm.mux.Lock()
  defer sm.mux.Unlock()
  if visited := sm.m[key]; visited {
    // already visited; return true
    return true
  } else {
    // not yet visited, add and return false
    sm.m[key] = true
    return false
  }
}


var fetched SafeMap = SafeMap{m: make(map[string]bool)}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  var wg sync.WaitGroup

  if depth <= 0 {
		return
	}

  if fetched.Found(url) {
    fmt.Printf("url %s already found.\n", url)
    return
  }

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
    wg.Add(1)
    go func(url string) {
      defer wg.Done()
      Crawl(url, depth-1, fetcher)
    }(u)
	}
  wg.Wait()
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
