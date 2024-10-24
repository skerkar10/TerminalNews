package ScrapePage

import (
	"github.com/gocolly/colly"
)

type Story struct {
  Name string
  Url string
  Source string
}

func ScrapeNews(collector *colly.Collector) []Story {
  var stories []Story

  collector.OnHTML("article", func(element *colly.HTMLElement) {
    currentStory := Story{}
    currentStory.Name = element.ChildText("a")
    currentStory.Url = "https://news.google.com" + element.ChildAttr("a", "href")

    stories = append(stories, currentStory)
  })

  collector.Visit("https://news.google.com")

  collector.Wait() //Avoid too many consecutive requests

  return stories
}
