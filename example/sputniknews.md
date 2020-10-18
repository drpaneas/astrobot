# Sputniknews


```go
package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url     string = "https://sputniknews.gr/astronomia/"
	baseURL string = "https://sputniknews.gr"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for Sputnik
var Doc *goquery.Document = getHTML(url)

// NewsDBsputnik db with the news
var NewsDBsputnik []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// GetNews fetches the news of sputniknews.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#rubric-major > div.b-stories.rubric-major > ul > li").Each(func(i int, s *goquery.Selection) {
		s.Find("div > div.b-stories__title > h2 > a").Each(func(i int, s *goquery.Selection) {
			link, ok = s.Attr("href")
			if ok {
				link = baseURL + link
				title = s.Text()
			}
		})
		s.Find("#rubric-major > div.b-stories.rubric-major > ul > li > a > picture > img").Each(func(i int, s *goquery.Selection) {
			img, ok := s.Attr("data-src")
			if ok {
				image = img
			}
		})
		s.Find("#rubric-major > div.b-stories.rubric-major > ul > li > div > div.b-stories__title > div > p").Each(func(i int, s *goquery.Selection) {
			desc = s.Text()
			// Remove newlines
			re := regexp.MustCompile(`\r?\n`)
			desc = re.ReplaceAllString(desc, " ")
			desc = stripSpaces(desc)
		})
		NewsDBsputnik = append(NewsDBsputnik, News{
			Description: desc,
			Image:       image,
			Link:        link,
			Title:       title,
			Source:      "sputniknews.gr",
		})
	})
}

func getHTML(page string) (doc *goquery.Document) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	standardClient := retryClient.StandardClient() // *http.Client

	// Request the HTML page.
	res, err := standardClient.Get(page)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func main() {
	GetNews()
	for _, v := range NewsDBsputnik {
		// fmt.Println(v.Title)
		fmt.Println(v.Description)
		// fmt.Println(v.Link)
		// fmt.Println(v.Image)
  // fmt.Println(v.Source)
	}
}
```