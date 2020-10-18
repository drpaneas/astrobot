package unboxholics

import (
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://unboxholics.com/tag/space"
)

// Doc for EarthSky
var Doc *goquery.Document = getHTML(url)

// NewsDBUnboxholics db with the news
var NewsDBUnboxholics []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// GetNews fetches the news of Unboxholics.com
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#main-container > div > div > article").Each(func(i int, s *goquery.Selection) {
		s.Find("div.entry__body.post-list__body.card__body > div.entry__header > h2 > a").Each(func(i int, s *goquery.Selection) {
			link, ok = s.Attr("href")
			if ok {
				title = s.Text()
			}
		})
		s.Find("div.entry__img-holder.post-list__img-holder.card__img-holder > img").Each(func(i int, s *goquery.Selection) {
			img, ok := s.Attr("data-src")
			if ok {
				image = img
			}
		})
		s.Find("div.entry__body.post-list__body.card__body > div.entry__excerpt > p").Each(func(i int, s *goquery.Selection) {
			desc = s.Text()
			// Remove newlines
			re := regexp.MustCompile(`\r?\n`)
			desc = re.ReplaceAllString(desc, " ")
		})
		NewsDBUnboxholics = append(NewsDBUnboxholics, News{
			Description: desc,
			Image:       image,
			Link:        link,
			Title:       title,
			Source:      "unboxholics.com",
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
