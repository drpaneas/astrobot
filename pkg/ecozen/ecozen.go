package ecozen

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
	url string = "https://ecozen.gr/category/epistimi/diastima/"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for ecozen.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBecozen db with the news
var NewsDBecozen []News

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

// titleQuery := "#news-item-209297 > header > h1"

// getIDs fetches the news of sputniknews.gr
func getIDs() []string {
	var id, idString string
	var ok bool

	Doc.Find("#content-container > div.cf > div.row.archive-page-container > div").Each(func(i int, s *goquery.Selection) {
		// Find the ID of the newspost
		s.Find("article").Each(func(i int, s *goquery.Selection) {
			id, ok = s.Attr("id")
			if ok {
				idString = idString + " " + id
			}
		})
	})
	if string(idString[0]) == " " {
		idString = trimFirstRune(idString)
	}
	idStringSlice := strings.Split(idString, " ")
	uniqueID := removeDuplicates(idStringSlice)
	return uniqueID
}

// GetNews gets the news for ecozen.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string

	linkQuery := fmt.Sprintf("#td-outer-wrap > div > div.td-container.td-category-container > div > div:nth-child(1) > div > div > div > div")
	Doc.Find(linkQuery).Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")

		linkQuery := fmt.Sprintf("#%s > div.td-big-grid-wrapper > div.td_module_mx5.td-animation-stack.td-big-grid-post-0.td-big-grid-post.td-big-thumb > div.td-module-thumb > a", id)
		Doc.Find(linkQuery).Each(func(i int, s *goquery.Selection) {
			tmpLink, ok := s.Attr("href")
			if ok {
				link = tmpLink
			}
		})

		imageQuery := fmt.Sprintf("#%s > div.td-big-grid-wrapper > div.td_module_mx5.td-animation-stack.td-big-grid-post-0.td-big-grid-post.td-big-thumb > div.td-module-thumb > a > img", id)
		Doc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
			image, _ = s.Attr("src")
		})

		doc := getHTML(link)
		descQuery := fmt.Sprintf("#gt-speech > p:nth-child(1)")
		doc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
			desc = s.Text()
			// Remove newlines
			re := regexp.MustCompile(`\r?\n`)
			desc = re.ReplaceAllString(desc, " ")
		})

		titleQuery := fmt.Sprintf("#%s > div.td-big-grid-wrapper > div.td_module_mx5.td-animation-stack.td-big-grid-post-0.td-big-grid-post.td-big-thumb > div.td-module-thumb > a > img", id)
		Doc.Find(titleQuery).Each(func(i int, s *goquery.Selection) {
			title, _ = s.Attr("title")
		})
		NewsDBecozen = append(NewsDBecozen, News{
			Description: desc,
			Image:       image,
			Link:        link,
			Source:      "ecozen.gr",
			Title:       title,
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
