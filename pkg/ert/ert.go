package ert

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://www.ert.gr/tag/diastima/"
)

// Doc for Ert
var Doc *goquery.Document = getHTML(url)

// NewsDBErt db with the news
var NewsDBErt []News

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
	Doc.Find("#td-outer-wrap > div.td-transition-content-and-menu.td-content-wrap > div:nth-child(2) > div > div > div.td-pb-span8.td-main-content > div > div").Each(func(i int, s *goquery.Selection) {
		s.Find("div.item-details > h3 > a").Each(func(i int, z *goquery.Selection) {
			link, ok = z.Attr("href")
			if ok {
				title = z.Text()
				s.Find("div.td-module-thumb > a > img").Each(func(i int, e *goquery.Selection) {
					img, ok := e.Attr("src")
					if ok {
						image = img
					}
				})
				s.Find("div.item-details > div.td-excerpt").Each(func(i int, w *goquery.Selection) {
					desc = stripSpaces(w.Text())
					desc = strings.Replace(desc, "\" ", "", -1)
					desc = strings.Replace(desc, " \"", "", -1)
				})
				NewsDBErt = append(NewsDBErt, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Title:       title,
					Source:      "ert.gr",
				})
			}
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
