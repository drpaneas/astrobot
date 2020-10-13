package tovima

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
	url string = "https://www.tovima.gr/tag/nasa"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for tovima.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBtovima db with the news
var NewsDBtovima []News

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

// GetNews gets the news for ecozen.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string

	linkQuery := fmt.Sprintf("#full-article-list > ul > li.modern-row.tablerow.mt_1 > div.mask-title.pos-rel > a")
	Doc.Find(linkQuery).Each(func(i int, s *goquery.Selection) {
		tmpLink, ok := s.Attr("href")
		if ok {
			link = tmpLink
			title = s.Text()
			if string(title[0]) == " " {
				title = trimFirstRune(title)
			}
			if strings.Contains(title, ":") {
				tmpTitle := strings.Split(title, ":")
				title = tmpTitle[1]
				if string(title[0]) == " " {
					title = trimFirstRune(title)
				}
			}
		}
	})

	doc := getHTML(link)
	imageQuery := fmt.Sprintf("div.hentry > div.article-main.tablerow.fullwidth.pos-rel.flex-container > div.mainpost > div.image-container > a > img")
	doc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
		image, _ = s.Attr("src")
	})

	descQuery := fmt.Sprintf("div.hentry > h2")
	doc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
		if string(desc[0]) == " " {
			desc = trimFirstRune(desc)
		}
	})
	NewsDBtovima = append(NewsDBtovima, News{
		Description: desc,
		Image:       image,
		Link:        link,
		Source:      "tovima.gr",
		Title:       title,
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
