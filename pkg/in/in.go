package in

import (
	"fmt"
	"log"
	"regexp"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://www.in.gr/tags/%ce%b4%ce%b9%ce%ac%cf%83%cf%84%ce%b7%ce%bc%ce%b1/"
)

// Doc for Ert
var Doc *goquery.Document = getHTML(url)

// NewsDBin db with the news
var NewsDBin []News

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

// Remove the first character of a string
func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// GetNews fetches the news of tanea.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool
	Doc.Find("#achive-page > div.category-slide > div > div > a.flow").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			s.Find("div > h3").Each(func(i int, s *goquery.Selection) {
				title = trimFirstRune(s.Text())
			})

			s.Find("div > p").Each(func(i int, s *goquery.Selection) {
				desc = trimFirstRune(s.Text())
			})
			s.Find("figure > img").Each(func(i int, s *goquery.Selection) {
				img, ok := s.Attr("data-src")
				if ok {
					image = img
				}
			})
			if len(NewsDBin) == 3 {
				return
			}
			NewsDBin = append(NewsDBin, News{
				Description: desc,
				Image:       image,
				Link:        link,
				Title:       title,
				Source:      "in.gr",
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
