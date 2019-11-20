package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func main() {
	getSS()
	readAll()
}

func getSS() {
	c := colly.NewCollector(
		//colly.Debugger(&debug.LogDebugger{}),
	)

	// Find and visit all links
	c.OnHTML(".context tbody tr", func(e *colly.HTMLElement) {
		s := new(SS)

		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			text := el.Text

			//if matched, _ := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}",text); matched {
			//	s.ip = text
			//}
			//if matched, _ := regexp.MatchString("([0-9]|[1-9]\\d{1,3}|[1-5]\\d{4}|6[0-4]\\d{4}|65[0-4]\\d{2}|655[0-2]\\d|6553[0-5])",text); matched {
			//	s.port = text
			//}

			if i == 1 {
				if matched, _ := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}",text); !matched {
					return
				}
				s.ip = text
			}
			if i == 2 {
				s.port = text
			}
			if i == 3 {
				s.password = text
			}
			if i == 4 {
				s.method = text
			}
		})
		fmt.Println(s)

		store(&s)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.youneed.win/free-ss")
}
