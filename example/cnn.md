# Code for CNN.gr

```go
package main

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
	url     string = "https://www.cnn.gr/tag/diastima"
	baseURL string = "https://www.cnn.gr"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for cnn.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBcnn db with the news
var NewsDBcnn []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	GreekTitle  string `json:"greektitle"`
	GreekDesc   string `json:"greekdesc"`
	Source      string `json:"source"`
	ID          string `json:"id"`
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
func getIDs() {
	var id, idString string
	var ok bool

	Doc.Find("#com_news > div.page > div.section.main-section.gap > div > div.news-items-section > div.flex-pack > div.flex-main > div.list-items > div").Each(func(i int, s *goquery.Selection) {
		// Find the ID of the newspost
		s.Find("a").Each(func(i int, s *goquery.Selection) {
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
	for _, v := range uniqueID {
		NewsDBcnn = append(NewsDBcnn, News{
			ID: v,
		})
	}
}

// GetNews gets the news for cnn.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	// Create the DB with IDs
	getIDs()

	// Loop through IDs and get data
	for i, v := range NewsDBcnn {
		id := strings.Replace(v.ID, "item-", "", -1)
		// Get the link
		Doc.Find("#item-" + id).Each(func(i int, s *goquery.Selection) {
			link, ok = s.Attr("href")
			if ok {
				link = baseURL + link
			}
		})

		// get the HTML per link
		doc := getHTML(link)
		titleQuery := fmt.Sprintf("#news-item-%s > header > h1", id)
		doc.Find(titleQuery).Each(func(i int, s *goquery.Selection) {
			title = s.Text()
		})

		imageQuery := fmt.Sprintf("#news-item-%s > div > div.flex-main > div > figure > picture > img", id)
		doc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
			image, _ = s.Attr("src")
		})

		descQuery := fmt.Sprintf("#news-item-%s > div > div.flex-main > div > div.main-content.story-content > div.main-intro.story-intro > p", id)
		doc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
			desc = s.Text()
			// Remove newlines
			re := regexp.MustCompile(`\r?\n`)
			desc = re.ReplaceAllString(desc, " ")
		})
		NewsDBcnn[i].Link = link
		NewsDBcnn[i].Title = title
		NewsDBcnn[i].Description = desc
		NewsDBcnn[i].Image = image
		NewsDBcnn[i].Source = "cnn.gr"
	}

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

func main() {
	GetNews()
	for _, v := range NewsDBcnn {
		fmt.Printf("ID    : %s\n", v.ID)
		fmt.Printf("Title : %s\n", v.Title)
		fmt.Printf("Descr : %s\n", v.Description)
		fmt.Printf("Link  : %s\n", v.Link)
		fmt.Printf("Image : %s\n", v.Image)
		fmt.Printf("Source: %s\n", v.Source)
	}
}
```