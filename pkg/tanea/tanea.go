package tanea

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://www.tanea.gr/tag/%ce%b4%ce%b9%ce%ac%cf%83%cf%84%ce%b7%ce%bc%ce%b1/"
)

// Doc for TaNea
var Doc *goquery.Document = getHTML(url)

// NewsDBTaNea db with the news
var NewsDBTaNea []News

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

	Doc.Find("#main > div > div.top-align.single-left-row.border-right > div.tablerow.fullwidth.fullheight > ul").Each(func(i int, s *goquery.Selection) {
		s.Find("li > div.mask-title > a").Each(func(i int, s *goquery.Selection) {
			link, ok = s.Attr("href")
			if ok {
				desc = s.Text()
				// Remove newlines
				re := regexp.MustCompile(`\r?\n`)
				desc = re.ReplaceAllString(desc, " ")
			}
		})
		s.Find("li > div.mask-image.pos-rel > a > div").Each(func(i int, s *goquery.Selection) {
			img, ok := s.Attr("style")
			if ok {
				tempAfter := strings.Split(img, "(")
				tempBefore := strings.Split(tempAfter[1], ")")
				image = tempBefore[0]
			}
		})
		s.Find("li > div.mask-title > a > span").Each(func(i int, s *goquery.Selection) {
			title = s.Text()
		})
		// Fix description
		desc = strings.ReplaceAll(desc, title, "")
		regex := regexp.MustCompile(`\r?\n`)
		desc = regex.ReplaceAllString(desc, " ")
		desc = strings.Replace(desc, "\t", "", -1)
		desc = stripSpaces(desc)
		desc = strings.Replace(desc, "\" ", "", -1)
		desc = strings.Replace(desc, " \"", "", -1)
		NewsDBTaNea = append(NewsDBTaNea, News{
			Description: desc,
			Image:       image,
			Link:        link,
			Title:       title,
			Source:      "tanea.gr",
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
