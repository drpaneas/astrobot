package news247

import (
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://www.news247.gr/diasthma"
)

// Doc for News247
var Doc *goquery.Document = getHTML(url)

// NewsDBNews247 db with the news
var NewsDBNews247 []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// GetNews fetches the news of NewsDBNews247.com
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("body > div.container > div.row.d-flex-row > div.col-xs-12.col-md-8 > div > section > div > article").Each(func(i int, s *goquery.Selection) {
		s.Find("h2 > a").Each(func(i int, s *goquery.Selection) {
			link, ok = s.Attr("href")
			if ok {
				title = s.Text()
			}
		})
		s.Find("figure > a > img").Each(func(i int, s *goquery.Selection) {
			img, ok := s.Attr("src")
			if ok {
				image = img
			}
		})
		s.Find("div > p").Each(func(i int, s *goquery.Selection) {
			desc = s.Text()
			// Remove newlines
			re := regexp.MustCompile(`\r?\n`)
			desc = re.ReplaceAllString(desc, " ")
		})
		NewsDBNews247 = append(NewsDBNews247, News{
			Description: desc,
			Image:       image,
			Link:        link,
			Title:       title,
			Source:      "news247.gr",
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
