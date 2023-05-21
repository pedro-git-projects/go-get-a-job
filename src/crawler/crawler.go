// Package crawler provides a simple web crawling functionality.
//
// The Crawler type represents a web crawler that can fetch data from a specified URL and apply
// optional filters to the crawled data. The Crawl method of the Crawler type utilizes goroutines
// for concurrent execution of the crawl operation. It sends a signal to indicate completion and
// waits for the operation to finish before returning.
//
// Example usage:
//
//	crawler := crawler.New()
//	crawler.Crawl("https://example.com", "filter1", "filter2")
//
// Note that the Crawl method uses a synchronization mechanism with the 'done' channel to ensure
// the main goroutine waits for the crawl operation to complete.
//
// For more information on the crawl operation and available filters, see the Crawl method documentation.
package crawler

type Crawler struct {
}

func New() *Crawler {
	return &Crawler{}
}

// Crawl fetches data from the specified URL and applies the provided filters to it.
//
// It creates a goroutine to concurrently execute the crawl operation. The crawl operation
// invokes the filterByTitle function with the given URL and filters. Once the crawl is
// complete, it sends a signal to the 'done' channel to indicate completion. The main
// goroutine waits for the signal from the 'done' channel before returning.
//
// Parameters:
//   - url: The URL to crawl and fetch data from.
//   - filters: Optional filters to apply to the crawled data.
//
// Example usage:
//
//	crawler := New()
//	crawler.Crawl("https://example.com", "filter1", "filter2")
//
// Note: The 'done' channel is used for synchronization, ensuring the main goroutine
// waits for the crawl operation to complete before returning.
func (*Crawler) Crawl(url string, filters ...string) {
	done := make(chan bool)
	go func() {
		filterByTitle(url, filters...)
		done <- true
	}()
	<-done
}
