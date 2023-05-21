package crawler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gitub.com/pedro-git-projects/go-get-a-job/src/parser"
	"golang.org/x/net/html"
)

// filterByTitle retrieves the HTML content from the specified URL and filters the job titles based on the provided titles.
// If no titles are provided, it filters the job titles containing the term "go".
// The filtered job titles are printed to the console.
func filterByTitle(url string, titles ...string) {
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

	parser.ProcessElements(doc, "h3", func(n *html.Node) {
		text := parser.GetText(n)
		if containsTitle(strings.ToLower(text), titles...) {
			fmt.Println("Job Title:", text)
		}
	})
}

// containsTitle checks if the specified text contains any of the provided titles.
// It returns true if any title is found, otherwise false.
func containsTitle(text string, titles ...string) bool {
	for _, title := range titles {
		if strings.Contains(text, title) {
			return true
		}
	}
	return false
}
