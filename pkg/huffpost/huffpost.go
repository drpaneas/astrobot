package huffpost

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
	url     string = "https://www.huffingtonpost.gr/news/astronomia/"
	baseURL string = "https://www.huffingtonpost.gr"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for huffingtonpost.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBhuffingtonpost db with the news
var NewsDBhuffingtonpost []News

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

	linkQuery := fmt.Sprintf("#card_center_1 > div > ul > li > h3 > a")
	Doc.Find(linkQuery).Each(func(i int, s *goquery.Selection) {
		tmpLink, ok := s.Attr("href")
		if ok {
			link = baseURL + tmpLink
		}
	})

	doc := getHTML(link)

	imageQuery := fmt.Sprintf("figure:nth-child(1) > span > img")
	doc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
		image, _ = s.Attr("src")
		tmp := strings.Split(image, "?")
		image = tmp[0]

	})

	descQuery := fmt.Sprintf("div.headline > h2")
	doc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
	})
	titleQuery := "div.headline > h1"
	doc.Find(titleQuery).Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})
	NewsDBhuffingtonpost = append(NewsDBhuffingtonpost, News{
		Description: desc,
		Image:       image,
		Link:        link,
		Source:      "huffingtonpost.gr",
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
