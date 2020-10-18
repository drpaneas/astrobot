package skai

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
	url       string = "https://www.skai.gr/tags/astronomia"
	baseURL   string = "https://www.skai.gr"
	urlSpaceX string = "https://www.skai.gr/tags/space-x"
)

func stripSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, " ")
	return fmt.Sprintf("%q", str)
}

// Doc for skai.gr
var Doc *goquery.Document = getHTML(url)
var docSpaceX *goquery.Document = getHTML(urlSpaceX)

// NewsDBskai db with the news
var NewsDBskai []News

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

	Doc.Find("body > div.dialog-off-canvas-main-canvas > main > div:nth-child(1) > div.categoryPinned.grid-x.grid-margin-x.medium-up-2.large-up-4 > div:nth-child(1) > div > div.cmnArticleTitlePad > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			link = baseURL + link
		}
		title = s.Text()
		if strings.Contains(title, "Αστρονομία") {
			title = strings.Split(title, "Αστρονομία")[1]
		}
		re := regexp.MustCompile(`\r?\n`)
		title = re.ReplaceAllString(title, " ")
		if string(title[0]) == " " {
			title = strings.TrimSpace(title)
		}
	})

	imageQuery := fmt.Sprintf("body > div.dialog-off-canvas-main-canvas > main > div:nth-child(1) > div.categoryPinned.grid-x.grid-margin-x.medium-up-2.large-up-4 > div:nth-child(1) > div > div.imgAligner > div > img")
	Doc.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
		tmpimage, _ := s.Attr("src")
		tmp := strings.Split(tmpimage, "?")
		image = baseURL + tmp[0]
	})
	doc := getHTML(link)
	descQuery := fmt.Sprintf("body > div.dialog-off-canvas-main-canvas > div.jscroll2.jscroll > div > main > div > div.viewWithSideBar > article > div.mainInfo > p")
	doc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
		if desc[0:1] == " " {
			desc = strings.TrimSpace(desc)
		}
	})
	NewsDBskai = append(NewsDBskai, News{
		Description: desc,
		Image:       image,
		Link:        link,
		Source:      "skai.gr",
		Title:       title,
	})

	docSpaceX.Find("body > div.dialog-off-canvas-main-canvas > main > div:nth-child(1) > div.categoryPinned.grid-x.grid-margin-x.medium-up-2.large-up-4 > div:nth-child(1) > div > div.cmnArticleTitlePad > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			link = baseURL + link
		}
		title = s.Text()
		if strings.Contains(title, "Space X") {
			title = strings.Split(title, "Space X")[1]
		}
		re := regexp.MustCompile(`\r?\n`)
		title = re.ReplaceAllString(title, " ")
		if string(title[0]) == " " {
			title = strings.TrimSpace(title)
		}
	})

	imageQuery = fmt.Sprintf("body > div.dialog-off-canvas-main-canvas > main > div:nth-child(1) > div.categoryPinned.grid-x.grid-margin-x.medium-up-2.large-up-4 > div:nth-child(1) > div > div.imgAligner > div > img")
	docSpaceX.Find(imageQuery).Each(func(i int, s *goquery.Selection) {
		tmpimage, _ := s.Attr("src")
		tmp := strings.Split(tmpimage, "?")
		image = baseURL + tmp[0]
	})
	docSpaceXdoc := getHTML(link)
	descQuery = fmt.Sprintf("body > div.dialog-off-canvas-main-canvas > div.jscroll2.jscroll > div > main > div > div.viewWithSideBar > article > div.mainInfo > p")
	docSpaceXdoc.Find(descQuery).Each(func(i int, s *goquery.Selection) {
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
		if desc[0:1] == " " {
			desc = strings.TrimSpace(desc)
		}
	})
	NewsDBskai = append(NewsDBskai, News{
		Description: desc,
		Image:       image,
		Link:        link,
		Source:      "skai.gr",
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
