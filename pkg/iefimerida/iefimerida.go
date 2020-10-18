package iefimerida

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
	url     string = "https://www.iefimerida.gr/tag/diastima"
	baseURL string = "https://www.iefimerida.gr"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for iefimerida.gr
var Doc *goquery.Document = getHTML(url)

// NewsDBiefimerida db with the news
var NewsDBiefimerida []News

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

// GetNews gets the news for cnn.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#taxonomy-term-180 > div.tag-content > div > article.iefimerida-article.teasers.big.big-teaser > h3 > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			link = baseURL + link
		}
		title = s.Text()
		re := regexp.MustCompile(`\r?\n`)
		title = re.ReplaceAllString(title, " ")
		if string(title[0]) == " " {
			title = strings.TrimSpace(title)
		}
	})
	descQuery := fmt.Sprintf("#taxonomy-term-180 > div.tag-content > div > article.iefimerida-article.teasers.big.big-teaser > div.field-summary")
	Doc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
		desc = trimFirstRune(desc)
	})
	imageQuery := fmt.Sprintf("#taxonomy-term-180 > div.tag-content > div > article.iefimerida-article.teasers.big.big-teaser > div.image-wrapper.tag-inside > a > picture > img")
	Doc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
		tmpimage, _ := s.Attr("src")
		tmp := strings.Split(tmpimage, "?")
		image = baseURL + tmp[0]
	})

	NewsDBiefimerida = append(NewsDBiefimerida, News{
		Description: desc,
		Image:       image,
		Link:        link,
		Source:      "iefimerida.gr",
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
