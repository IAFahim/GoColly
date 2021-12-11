package main

import (
	"fmt"
	"net/url"
)

import (
	"github.com/gocolly/colly/v2"
)

func WebScraper(link *url.URL) {
	c := colly.NewCollector(
		colly.AllowedDomains(link.Host),
	)

	c.OnHTML(".vector-body p", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(link.String())

}

func main() {
	u, err := url.Parse("https://en.wikipedia.org/wiki/Web_scraping")
	if err == nil {
		println(err)
	}
	WebScraper(u)
}
