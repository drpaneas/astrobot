package pontosnews

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
	url     string = "http://www.pontos-news.gr/tags/astronomia"
	baseURL string = "http://www.pontos-news.gr"
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

// Doc for pontos-news.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBpontos db with the news
var NewsDBpontos []News

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
	for _, value := range NewsDBpontos {
		if link == value.Link {
			return true
		}
	}
	return false
}

// GetNews gets the news for cnn.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#block-system-main > div > div > div.view.view-taxonomy-term.view-id-taxonomy_term.view-display-id-page.main-teaser.view-dom-id- > div > div.views-row.views-row-1.views-row-odd.views-row-first > div > div > span > div > div.content > div.inner > h4 > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			if !testURLReachable(link) {
				link = baseURL + link
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				if string(title[:0]) == " " {
					title = strings.TrimSpace(title)
				}
			}

			// Image:
			imageQuery := fmt.Sprintf("#block-system-main > div > div > div.view.view-taxonomy-term.view-id-taxonomy_term.view-display-id-page.main-teaser.view-dom-id- > div > div.views-row.views-row-1.views-row-odd.views-row-first > div > div > span > div > div.content > div.image > a > img")
			Doc.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
				tmpImage, imageExists := w.Attr("src")
				if imageExists {
					tmp := strings.Split(tmpImage, "?")
					if !testURLReachable(tmpImage) {
						image = baseURL + tmp[0]
					} else {
						image = tmp[0]
					}
				}
			})

			// Description
			descQuery := fmt.Sprintf("#block-system-main > div > div > div.view.view-taxonomy-term.view-id-taxonomy_term.view-display-id-page.main-teaser.view-dom-id- > div > div.views-row.views-row-1.views-row-odd.views-row-first > div > div > span > div > div.content > div.inner > div.body")
			Doc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
				desc = u.Text()
				desc = strings.TrimSpace(desc)
			})

			NewsDBpontos = append(NewsDBpontos, News{
				Description: desc,
				Image:       image,
				Link:        link,
				Source:      "pontos-news.gr",
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
