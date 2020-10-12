# Contribute

If you want to add a new website to parse for news, send me PR into this directory parsing it.
Here's an example:

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
	url string = "https://www.space.com/news"
)

// Doc for Space.com
var Doc *goquery.Document = getHTML(url)

// NewsDBSpace db with the news
var NewsDBSpace []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	GreekTitle  string `json:"greektitle"`
	GreekDesc   string `json:"greekdesc"`
	Source      string `json:"source"`
}

// GetNews fetches the news of Space.com
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#content > section > section > div.listingResults.mixed > div > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {

			s.Find("article > div.content > header > h3").Each(func(i int, s *goquery.Selection) {
				title = s.Text()
			})
			s.Find("article > div.image > figure").Each(func(i int, s *goquery.Selection) {
				img, ok := s.Attr("data-original")
				if ok {
					image = img
				}
			})
			s.Find("article > div.content > p").Each(func(i int, s *goquery.Selection) {
				desc = s.Text()
				// Remove newlines
				re := regexp.MustCompile(`\r?\n`)
				desc = re.ReplaceAllString(desc, " ")
				// Remove first char if it's empty
				if desc[:0] == "" {
					desc = desc[1:]
				}
			})
			NewsDBSpace = append(NewsDBSpace, News{
				Description: desc,
				Image:       image,
				Link:        link,
				Title:       title,
				Source:      "space.com",
			})
		}
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
	 for _, v := range NewsDBSpace {
		// fmt.Fprintf("Title:%v\nDescri:%v\nLink:%v\nImage:%v\nSource:%v\n\n", v.Title, v.Description, v.Link, v.Image, v.Source)
        fmt.Println(v.Title)
        fmt.Println(v.Description)
        fmt.Println(v.Link)
        fmt.Println(v.Image)
        fmt.Println(v.Source)
	}
}
```

If this is working, then I will take your code and embedded it.
