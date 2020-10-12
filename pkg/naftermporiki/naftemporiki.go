package naftermporiki

import (
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url     string = "https://www.naftemporiki.gr/tag/96/diastima"
	baseURL string = "https://www.naftemporiki.gr"
)

// Doc for Naftermporiki
var Doc *goquery.Document = getHTML(url)

// NewsDBNaftermporiki db with the news
var NewsDBNaftermporiki []News

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

// GetNews fetches the news of Naftermporiki.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#leftArea > div.latest > div.topic > ul").Each(func(i int, s *goquery.Selection) {
		s.Find("li > div.summary > h4 > a").Each(func(i int, s *goquery.Selection) {
			link, ok = s.Attr("href")
			if ok {
				link = baseURL + link
				title = s.Text()
			}
		})
		s.Find("li > div.photoContainer > a > img").Each(func(i int, s *goquery.Selection) {
			img, ok := s.Attr("src")
			if ok {
				image = baseURL + img
			}
		})
		s.Find("li > div.summary > div > p").Each(func(i int, s *goquery.Selection) {
			desc = s.Text()
			// Remove newlines
			re := regexp.MustCompile(`\r?\n`)
			desc = re.ReplaceAllString(desc, " ")
		})
		NewsDBNaftermporiki = append(NewsDBNaftermporiki, News{
			Description: desc,
			Image:       image,
			Link:        link,
			Title:       title,
			Source:      "naftemporiki.gr",
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
