package gazzetta

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
	url     string = "https://www.gazzetta.gr/tag/diastima"
	baseURL string = "https://www.gazzetta.gr"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for gazzetta.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBgazzetta db with the news
var NewsDBgazzetta []News

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

// Remove the duplicate elements from slice
func removeDuplicates(elements []string) []string { // change string to int here if required
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{} // change string to int here if required
	result := []string{}             // change string to int here if required

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// GetNews gets the news for ecozen.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string

	linkQuery := fmt.Sprintf("#block-system-main > div > div > div > div.view-content.clearfix > div.views-row.views-row-1.views-row-odd.views-row-first > div > div > div.field-name-title > h2 > a")
	Doc.Find(linkQuery).Each(func(i int, s *goquery.Selection) {
		tmpLink, ok := s.Attr("href")
		if ok {
			link = baseURL + tmpLink
		}
	})

	doc := getHTML(link)
	imageQuery := fmt.Sprintf("#block-system-main > div > div > div > div.group-header.clearfix > div.field.field-name-full-screen-cover-image.field-type-ds.field-label-hidden > div > div > figure > img")
	doc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
		tmpimage, _ := s.Attr("src")
		tmp := strings.Split(tmpimage, "?")
		image = baseURL + tmp[0]
	})

	descQuery := fmt.Sprintf("#block-system-main > div > div > div > div.group-left > div.field.field-name-article-body-with-ads.field-type-ds.field-label-hidden > div > div > p:nth-child(1)")
	doc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
	})

	titleQuery := "#block-system-main > div > div > div > div.group-header.clearfix > div.field.field-name-title.field-type-ds.field-label-hidden > div > div > h1"
	doc.Find(titleQuery).Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})
	NewsDBgazzetta = append(NewsDBgazzetta, News{
		Description: desc,
		Image:       image,
		Link:        link,
		Source:      "gazzetta.gr",
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
