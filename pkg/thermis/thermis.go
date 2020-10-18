package thermis

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
	url     string = "https://www.thermisnews.gr/search/label/%CE%94%CE%99%CE%91%CE%A3%CE%A4%CE%97%CE%9C%CE%91"
	baseURL string = "https://www.thermisnews.gr"
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

// Doc for thermisnews.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBthermis db with the news
var NewsDBthermis []News

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
	for _, value := range NewsDBthermis {
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

	Doc.Find("#Blog1 > div.blog-posts.hfeed > div:nth-child(1) > div > article > font > h2 > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			if !testURLReachable(link) {
				link = baseURL + link
			}

			// TItle:
			titleQuery := fmt.Sprintf("#Blog1 > div.blog-posts.hfeed > div:nth-child(1) > div > article > font > h2")
			Doc.Find(titleQuery).Each(func(j int, w *goquery.Selection) {
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				title = strings.TrimSpace(title)
			})

			// Image:
			imageQuery := fmt.Sprintf("#Blog1 > div.blog-posts.hfeed > div:nth-child(1) > div > div.block-image > div.thumb > a")
			Doc.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
				tmpImage, imageExists := w.Attr("style")
				if imageExists {
					tmpRemoveLeft := strings.Split(tmpImage, "(")
					tmp := strings.Split(tmpRemoveLeft[1], ")")
					if !testURLReachable(tmpImage) {
						image = baseURL + tmp[0]
					} else {
						image = tmp[0]
					}
				}
			})

			// Description
			descQuery := fmt.Sprintf("#Blog1 > div.blog-posts.hfeed > div:nth-child(1) > div > article > div > div.resumo > span")
			Doc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
				desc = u.Text()
				desc = strings.TrimSpace(desc)
			})

			NewsDBthermis = append(NewsDBthermis, News{
				Description: desc,
				Image:       image,
				Link:        link,
				Source:      "thermisnews.gr",
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
