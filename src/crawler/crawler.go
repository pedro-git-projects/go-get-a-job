package crawler

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"gitub.com/pedro-git-projects/go-get-a-job/src/parser"
	"golang.org/x/net/html"
)

type Crawler struct {
	mu sync.Mutex
}

func New() *Crawler {
	return &Crawler{}
}

func mkreq(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr,
			"request failed with status code %d\n",
			res.StatusCode)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	parser.ProcessElements(doc, "h1", func(n *html.Node) {
		text := parser.GetText(n)
		fmt.Printf("Heading: %s\n", text)
	})
}

func (*Crawler) Crawl(urls ...string) {
	for _, url := range urls {
		mkreq(url)
	}
}
