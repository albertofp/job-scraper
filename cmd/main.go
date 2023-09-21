package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Job struct {
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Remote      bool     `json:"remote"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"created_at"`
	JobTypes    []string `json:"job_types"`
}

func main() {
	// jobs := make([]Job, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("de.indeed.com"),
		colly.AllowURLRevisit(),
	)

	extensions.RandomUserAgent(c)

	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*indeed.*",
		Parallelism: 3,
		RandomDelay: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "\nError:", err)
	})

	c.OnHTML("h2.jobTitle > a > span", func(e *colly.HTMLElement) {
		fmt.Println(e.Text, "found")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c.Visit("https://de.indeed.com/jobs?q=Software+Engineer&l=berlin&vjk=1a4c986f12787f88"); err != nil {
		panic(err)
	}
}
