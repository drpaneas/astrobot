package physicsgg

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
	url     string = "https://physicsgg.me/category/%ce%b1%cf%83%cf%84%cf%81%ce%bf%cf%86%cf%85%cf%83%ce%b9%ce%ba%ce%b7/"
	baseURL string = "https://physicsgg.me"
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

// Doc for physicsgg.me
var Doc *goquery.Document = getHTML(url)

// NewsDBphysicsgg db with the news
var NewsDBphysicsgg []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

func fixTitle(title string) string {
	if strings.Contains(title, "Επιστήμη") {
		title = strings.Split(title, "Επιστήμη")[1]
	}
	if strings.Contains(title, "Διάστημα") {
		title = strings.Split(title, "Διάστημα")[1]
	}
	if strings.Contains(title, "NASA") {
		title = strings.Split(title, "NASA")[1]
	}
	if strings.Contains(title, "Space X") {
		title = strings.Split(title, "Space X")[1]
	}
	if strings.Contains(title, "Αστρονομία") {
		title = strings.Split(title, "Αστρονομία")[1]
	}
	return title
}

// Remove the first character of a string
func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func thisNewExistsInAnotherTag(link string) bool {
	for _, value := range NewsDBphysicsgg {
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

	Doc.Find("#post-wrapper > div > h2 > a").Each(func(i int, s *goquery.Selection) {
		if len(NewsDBphysicsgg) <= 3 { // limit to 3 articles

			link, ok = s.Attr("href")
			if ok {
				// Title:
				title = s.Text()
				title = fixTitle(title)
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				if string(title[:0]) == " " {
					title = strings.TrimSpace(title)
				}

				individualDoc := getHTML(link)

				// Image:
				individualDoc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
					href, _ := item.Attr("href")
					if strings.Contains(href, ".png") {
						image = href
					} else if strings.Contains(href, ".jpg") {
						image = href
					}
				})

				// Description
				descQuery := fmt.Sprintf("div.entry > p:nth-child(3)")
				individualDoc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
					desc = s.Text()
				})

				NewsDBphysicsgg = append(NewsDBphysicsgg, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Source:      "physicsgg.me",
					Title:       title,
				})
			}
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
