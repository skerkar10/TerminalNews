package main

import (
	"TerminalNews/ScrapePage"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)


func main() {
  webScraper := colly.NewCollector(
    // colly.AllowedDomains("www.news.google.com"),
    colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
    colly.Async(true),
  )

  extensions.RandomUserAgent(webScraper)
  extensions.Referer(webScraper)

  stories := ScrapePage.ScrapeNews(webScraper)

  for _, story := range stories {
    fmt.Printf("Title: %s\nURL: %s\n\n", story.Name, story.Url)
  }
}
