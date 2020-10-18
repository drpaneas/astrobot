package protothema

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://www.protothema.gr/tag/diastima/"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for ProtoThema
var Doc *goquery.Document = getHTML(url)

// NewsDBProtoThema db with the news
var NewsDBProtoThema []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// GetNews fetches the news of protothema.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("body > div.outer > main > section > div > div.mainWrp > div > div > div > div").Each(func(i int, s *goquery.Selection) {
		s.Find("div > article > div > div.heading > h3 > a").Each(func(i int, s *goquery.Selection) {
			link, ok = s.Attr("href")
			if ok {
				title = s.Text()
			}
		})
		s.Find("div > article > figure > a > picture > img").Each(func(i int, s *goquery.Selection) {
			img, ok := s.Attr("data-src")
			if ok {
				image = img
			}
		})
		s.Find("div > article > div > div.txt > p").Each(func(i int, s *goquery.Selection) {
			desc = s.Text()
			// Remove newlines
			re := regexp.MustCompile(`\r?\n`)
			desc = re.ReplaceAllString(desc, " ")
			desc = stripSpaces(desc)
		})
		NewsDBProtoThema = append(NewsDBProtoThema, News{
			Description: desc,
			Image:       image,
			Link:        link,
			Title:       title,
			Source:      "protothema.gr",
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
