package astronio

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
	"log"
	"regexp"
	"strings"
)

const (
	url     string = "https://www.astronio.gr/archives/category/astronews"
	baseURL string = "https://www.astronio.gr/"
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

// Doc for astronio.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBastronio db with the news
var NewsDBastronio []News

// News represent a new article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// GetNews gets the news for astronio.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#main > div > ul > li > article > div > a").Each(func(i int, s *goquery.Selection) {
		// Get the link
		link, ok = s.Attr("href")
		if ok {
			if !testURLReachable(link) {
				link = baseURL + link
			}

			// From the link, get the Thumbnail
			image, _ = s.Attr("data-bgset")

			// Update the Astronews News DB
			NewsDBastronio = append(NewsDBastronio, News{
				Link:        link,
				Source:      "astronio.gr",
			})
		}
	})

	// For every item in the Astronews News DB, get the Title and the first paragraph for the desc.
	for i, article := range NewsDBastronio {
		fmt.Println("Parsing article:", article.Link)

		// Find the article number
		number := article.Link[strings.LastIndex(article.Link, "/")+1:]
		doc := getHTML(article.Link)

		// Get Title
		titleQuery := fmt.Sprintf("#post-%s > div.header-standard.header-classic.single-header > h1", number)
		doc.Find(titleQuery).Each(func(j int, s *goquery.Selection) {
			title = s.Text()
			re := regexp.MustCompile(`\r?\n`)
			title = re.ReplaceAllString(title, " ")
			title = strings.TrimSpace(title)
			title = replaceGreekRunes(title)
			NewsDBastronio[i].Title = title
		})

		//Get Description
		descQuery := fmt.Sprintf("#penci-post-entry-inner > p:nth-child(1)")
		doc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
			desc = u.Text()
			desc = strings.TrimSpace(desc)
			desc = replaceGreekRunes(desc)
			NewsDBastronio[i].Description = desc
		})

		if NewsDBastronio[i].Description == "" {
			descQuery = fmt.Sprintf("#penci-post-entry-inner > p:nth-child(2)")
			doc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
				desc = u.Text()
				desc = strings.TrimSpace(desc)
				desc = replaceGreekRunes(desc)
				NewsDBastronio[i].Description = desc
			})
		}

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


func replaceGreekRunes(title string) string {
	if strings.Contains(title, "“") {
		title = strings.ReplaceAll(title, "“", "")
	}
	if strings.Contains(title, "”") {
		title = strings.ReplaceAll(title, "”", "")
	}
	if strings.Contains(title, "\"") {
		fmt.Println("\"", title)
		title = strings.ReplaceAll(title, "\"", "")
	}
	if strings.Contains(title, "/") {
		title = strings.ReplaceAll(title, "/", "")
	}
	if strings.Contains(title, "\\") {
		title = strings.ReplaceAll(title, "\\", "")
	}
	if strings.Contains(title, "«") {
		title = strings.ReplaceAll(title, "«", "")
	}
	if strings.Contains(title, "»") {
		title = strings.ReplaceAll(title, "»", "")
	}
	if strings.Contains(title, "#") {
		title = strings.ReplaceAll(title, "#", "")
	}
	if strings.Contains(title, ":") {
		title = strings.ReplaceAll(title, ":", "")
	}
	return title
}
