package universetoday

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url string = "https://www.universetoday.com/"
)

// Doc for universetoday
var Doc *goquery.Document = getHTML(url)

// NewsDBUniverseToday db with the news
var NewsDBUniverseToday []News

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

func getImage(postid string) (image string) {
	query := fmt.Sprintf("#%s > div.post-thumbnail > a > img", postid)
	Doc.Find(query).Each(func(i int, s *goquery.Selection) {
		image, _ = s.Attr("src")
	})
	return image
}

func getTitleAndLink(postid string) (title, link string) {
	query := fmt.Sprintf("#%s > header > h3 > a", postid)
	Doc.Find(query).Each(func(i int, s *goquery.Selection) {
		title = s.Text()
		link, _ = s.Attr("href")
	})
	return title, link
}

func getDescription(postid string) (desc string) {
	query := fmt.Sprintf("#%s > div.entry-content", postid)
	Doc.Find(query).Each(func(i int, s *goquery.Selection) {
		// desc, descExists := s.Attr("p")
		// if descExists {
		// 	fmt.Println(desc)
		// }
		desc = s.Text()
		// Remove newlines
		re := regexp.MustCompile(`\r?\n`)
		desc = re.ReplaceAllString(desc, " ")
		// Remove first char if it's empty
		if desc[:0] == "" {
			desc = desc[1:]
		}
	})
	return desc
}

// GetNews fetches the news of UniverseToday
func GetNews() {
	Doc.Find("#main").Each(func(i int, s *goquery.Selection) {
		s.Find("article").Each(func(i int, s *goquery.Selection) {
			postID, newsPostExists := s.Attr("id")
			if newsPostExists {
				title, link := getTitleAndLink(postID)
				image := getImage(postID)
				desc := getDescription(postID)
				NewsDBUniverseToday = append(NewsDBUniverseToday, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Title:       title,
					Source:      "universetoday.com",
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
