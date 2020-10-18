package earthsky

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://earthsky.org/space"
)

// Doc for EarthSky
var Doc *goquery.Document = getHTML(url)

// NewsDBEarthSky db with the news
var NewsDBEarthSky []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

func getImage(postid string) (image string) {
	query := fmt.Sprintf("#%s > div > div.post_archive_content.span-11.last > div.entry-content.entry-summary > a > div", postid)
	Doc.Find(query).Each(func(i int, s *goquery.Selection) {
		kati, _ := s.Attr("style")
		katiAllo := strings.Split(kati, "'")
		image = katiAllo[1]
	})
	return image
}
func getTitleAndLink(postid string) (title, link string) {
	query := fmt.Sprintf("#%s > div > div.post_archive_content.span-11.last > div.entry_header > h2 > a", postid)
	Doc.Find(query).Each(func(i int, s *goquery.Selection) {
		title = s.Text()
		link, _ = s.Attr("href")
	})
	return title, link
}
func getDescription(postid string) (desc string) {
	query := fmt.Sprintf("#%s > div > div.post_archive_content.span-11.last > div.entry-content.entry-summary > p", postid)
	Doc.Find(query).Each(func(i int, s *goquery.Selection) {
		// desc, descExists := s.Attr("p")
		// if descExists {
		// 	fmt.Println(desc)
		// }
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
	})
	return desc
}

// GetNews fetches the news of UniverseToday
func GetNews() {
	Doc.Find("#content").Each(func(i int, s *goquery.Selection) {
		s.Find("div").Each(func(i int, s *goquery.Selection) {
			postID, newsPostExists := s.Attr("id")
			if newsPostExists {
				title, link := getTitleAndLink(postID)
				image := getImage(postID)
				desc := getDescription(postID)
				NewsDBEarthSky = append(NewsDBEarthSky, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Title:       title,
					Source:      "earthsky.org",
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
