package main

import "gitub.com/pedro-git-projects/go-get-a-job/src/crawler"

func main() {
	crawler.New().Crawl(crawler.Seeds[0], "go", "golang")
}
