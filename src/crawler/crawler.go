package crawler

import (
	"sync"
)

type Crawler struct {
	mu sync.Mutex
}

func New() *Crawler {
	return &Crawler{}
}

func (*Crawler) Crawl(url string, filters ...string) {
	filterByTitle(url, filters...)
}
