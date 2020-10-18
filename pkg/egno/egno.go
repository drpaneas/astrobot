package egno

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url     string = "https://egno.gr/category/space/"
	baseURL string = "https://egno.gr"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

func testURLReachable(link string) bool {
	if strings.Contains(link, "https") {
		return true
	}
	return false
}

// Doc for egno.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBegno db with the news
var NewsDBegno []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// Remove the first character of a string
func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func thisNewExistsInAnotherTag(link string) bool {
	for _, value := range NewsDBegno {
		if link == value.Link {
			return true
		}
	}
	return false
}

// GetNews gets the news for thermisnews.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#home-left-big > div.home-widget > ul > li:nth-child(1) > div.category3-text > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			if !testURLReachable(link) {
				link = baseURL + link
			}

			// TItle:
			titleQuery := fmt.Sprintf("#home-left-big > div.home-widget > ul > li:nth-child(1) > div.category3-text > a")
			Doc.Find(titleQuery).Each(func(j int, w *goquery.Selection) {
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				title = strings.TrimSpace(title)
			})

			// Image:
			imageQuery := fmt.Sprintf("#home-left-big > div.home-widget > ul > li:nth-child(1) > div.category3-image > a > img")
			Doc.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
				tmpImage, imageExists := w.Attr("src")
				if imageExists {
					if !testURLReachable(tmpImage) {
						image = baseURL + tmpImage
					} else {
						image = tmpImage
					}
				}
			})

			// Description
			descQuery := fmt.Sprintf("#home-left-big > div.home-widget > ul > li:nth-child(1) > div.category3-text > p")
			Doc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
				desc = u.Text()
				desc = strings.TrimSpace(desc)
			})

			NewsDBegno = append(NewsDBegno, News{
				Description: desc,
				Image:       image,
				Link:        link,
				Source:      "egno.gr",
				Title:       title,
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
