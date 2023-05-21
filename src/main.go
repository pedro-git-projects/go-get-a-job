package main

import (
	"sync"

	"gitub.com/pedro-git-projects/go-get-a-job/src/crawler"
)

func main() {
	c := crawler.New()
	var wg sync.WaitGroup

	crawlCalls := []struct {
		url     string
		filters []string
	}{
		{url: crawler.Seeds[0], filters: []string{"go", "golang"}},
		{url: crawler.Seeds[1], filters: []string{"go", "golang"}},
		{url: crawler.Seeds[2], filters: []string{"estágio", "estagio"}},
		{url: crawler.Seeds[2], filters: []string{"java"}},
		{url: crawler.Seeds[3], filters: []string{"java"}},
	}

	for _, call := range crawlCalls {
		wg.Add(1)
		go func(url string, filters []string) {
			defer wg.Done()
			c.Crawl(url, filters...)
		}(call.url, call.filters)
	}
	wg.Wait()
}
