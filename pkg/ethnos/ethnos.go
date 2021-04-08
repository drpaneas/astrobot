package ethnos

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	url           string = "https://ethnos.gr/diastima"
	urlAstronomia string = "https://ethnos.gr/astronomia"
	baseURL       string = "https://ethnos.gr"
	urlNASA       string = "https://ethnos.gr/nasa"
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

// Doc for ethnos.gr
var Doc *goquery.Document = getHTML(url)
var docAstronomia *goquery.Document = getHTML(urlAstronomia)
var docNASA *goquery.Document = getHTML(urlNASA)

// NewsDBethnos db with the news
var NewsDBethnos []News

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

func thisNewExistsInAnotherTag(link string) bool {
	for _, value := range NewsDBethnos {
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

	Doc.Find("#page > div > main > section.main__full > section > div > article > a.full-link").Each(func(i int, s *goquery.Selection) {
		if len(NewsDBethnos) <= 3 { // limit to 3 articles

			link, ok = s.Attr("href")
			if ok {
				// Title:
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				if string(title[:0]) == " " {
					title = strings.TrimSpace(title)
				}

				if !testURLReachable(link) {
					link = baseURL + link
				}

				individualDoc := getHTML(link)

				// Image:
				imageQuery := fmt.Sprintf("#page > div > main > section.main__content > article > figure > img")
				individualDoc.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
					tmpImage, imageExists := w.Attr("src")
					if imageExists {
						tmp := strings.Split(tmpImage, "?")
						if !testURLReachable(tmpImage) {
							image = baseURL + tmp[0]
						} else {
							image = tmp[0]
						}
					}
				})

				// Description
				descQuery := fmt.Sprintf("#page > div > main > section.main__content > article > section > div.article__main > div.article__summary")
				individualDoc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
					desc = u.Text()
					desc = strings.TrimSpace(desc)
				})

				NewsDBethnos = append(NewsDBethnos, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Source:      "ethnos.gr",
					Title:       title,
				})
			}
		}
	})

	// New Tag Astronomia
	docAstronomia.Find("#page > div > main > section.main__full > section > div > article > a.full-link").Each(func(i int, s *goquery.Selection) {
		if len(NewsDBethnos) <= 6 { // limit to 6 articles (3 from this tag + 3 from the other tag)

			link, ok = s.Attr("href")
			if ok {
				// Title:
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				if string(title[:0]) == " " {
					title = strings.TrimSpace(title)
				}

				if !testURLReachable(link) {
					link = baseURL + link
				}

				individualDoc := getHTML(link)

				// Image:
				imageQuery := fmt.Sprintf("#page > div > main > section.main__content > article > figure > img")
				individualDoc.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
					tmpImage, imageExists := w.Attr("src")
					if imageExists {
						tmp := strings.Split(tmpImage, "?")
						if !testURLReachable(tmpImage) {
							image = baseURL + tmp[0]
						} else {
							image = tmp[0]
						}
					}
				})

				// Description
				descQuery := fmt.Sprintf("#page > div > main > section.main__content > article > section > div.article__main > div.article__summary")
				individualDoc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
					desc = u.Text()
					desc = strings.TrimSpace(desc)
				})

				if !thisNewExistsInAnotherTag(link) {
					NewsDBethnos = append(NewsDBethnos, News{
						Description: desc,
						Image:       image,
						Link:        link,
						Source:      "ethnos.gr",
						Title:       title,
					})
				}
			}
		}
	})

	// New Tag NASA
	// New Tag Astronomia
	docAstronomia.Find("#page > div > main > section.main__full > section > div > article > a.full-link").Each(func(i int, s *goquery.Selection) {
		if len(NewsDBethnos) <= 9 { // limit to 6 articles (3 from this tag + 3 from the other tag + 3 from the first tag)

			link, ok = s.Attr("href")
			if ok {
				// Title:
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				if string(title[:0]) == " " {
					title = strings.TrimSpace(title)
				}

				if !testURLReachable(link) {
					link = baseURL + link
				}

				individualDoc := getHTML(link)

				// Image:
				imageQuery := fmt.Sprintf("#page > div > main > section.main__content > article > figure > img")
				individualDoc.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
					tmpImage, imageExists := w.Attr("src")
					if imageExists {
						tmp := strings.Split(tmpImage, "?")
						if !testURLReachable(tmpImage) {
							image = baseURL + tmp[0]
						} else {
							image = tmp[0]
						}
					}
				})

				// Description
				descQuery := fmt.Sprintf("#page > div > main > section.main__content > article > section > div.article__main > div.article__summary")
				individualDoc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
					desc = u.Text()
					desc = strings.TrimSpace(desc)
				})

				if !thisNewExistsInAnotherTag(link) {
					NewsDBethnos = append(NewsDBethnos, News{
						Description: desc,
						Image:       image,
						Link:        link,
						Source:      "ethnos.gr",
						Title:       title,
					})
				}
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
		fmt.Errorf("couldn't load the page %s, because of error: %q", page, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Errorf("cannot load HTML doc: %q", err)
	}
	return doc
}
