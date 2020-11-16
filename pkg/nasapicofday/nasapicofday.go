package nasapicofday

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url              string = "https://apod.nasa.gov/apod/archivepix.html"
	baseURL          string = "https://apod.nasa.gov/apod/"
	numberOfPictures int    = 1
)

func testURLReachable(link string) bool {
	if strings.Contains(link, "https://") {
		return true
	}
	return false
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

// Doc for nasa picture
var Doc *goquery.Document = getHTML(url)

// NewsDBNasaImage db with the news
var NewsDBNasaImage []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// GetNews gets the news for Picture of the Day @ Nasa
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	titleQuery := fmt.Sprintf("body > b > a")
	Doc.Find(titleQuery).Each(func(i int, s *goquery.Selection) {
		title = s.Text()
		url, ok := s.Attr("href")
		if ok {
			link = url
			if !testURLReachable(link) {
				link = baseURL + url
			}
			if len(NewsDBNasaImage) < numberOfPictures {
				SpecificDoc := getHTML(link)
				imageQuery := fmt.Sprintf("body > center > p > a > img")
				SpecificDoc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
					image, _ = s.Attr("src")
					if !testURLReachable(image) {
						image = baseURL + image
					}
				})
				counter := 0
				descQuery := fmt.Sprintf("body p")
				SpecificDoc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
					counter = counter + 1
					if counter == 3 {
						desc = s.Text()
						desc = strings.Replace(desc, "\n", "", -1) // remove line breaks
						desc = strip.StripTags(desc)               // remove html tags if any
						if strings.Contains(desc, "Explanation:") {
							tmp := strings.SplitAfter(desc, "Explanation:")
							desc = tmp[1]
						}
						// Remove trailing whitespace at the beginning
						for {
							if strings.Contains(desc[0:1], " ") {
								desc = trimFirstRune(desc)
							} else {
								break
							}
						}
					}
				})

				NewsDBNasaImage = append(NewsDBNasaImage, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Source:      "nasa.gov",
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
