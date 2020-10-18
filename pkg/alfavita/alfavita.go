package alfavita

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
	url     string = "https://www.alfavita.gr/diastima"
	urlNASA string = "https://www.alfavita.gr/nasa"
	baseURL string = "https://www.alfavita.gr"
	urlMoon string = "https://www.alfavita.gr/selini"
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

// Doc for alfavita.gr
var Doc *goquery.Document = getHTML(url)
var docNasa *goquery.Document = getHTML(urlNASA)
var docMoon *goquery.Document = getHTML(urlMoon)

// NewsDBalfavita db with the news
var NewsDBalfavita []News

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
	for _, value := range NewsDBalfavita {
		if link == value.Link {
			return true
		}
	}
	return false
}

// GetNews gets the news for thermisnews.gr
func GetNews() {
	var title string
	var image string
	var desc string
	var link string
	var ok bool

	Doc.Find("#page > main > section.main__content > div.list-promo > article:nth-child(1) > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			if !testURLReachable(link) {
				link = baseURL + link
			}

			// TItle:
			titleQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > div.list-promo__title")
			Doc.Find(titleQuery).Each(func(j int, w *goquery.Selection) {
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				title = strings.TrimSpace(title)
			})

			// Image:
			imageQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > figure > img")
			Doc.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
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
			descQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > div.list-promo__summary")
			Doc.Find(descQuery).Each(func(p int, u *goquery.Selection) {
				desc = u.Text()
				desc = strings.TrimSpace(desc)
			})

			NewsDBalfavita = append(NewsDBalfavita, News{
				Description: desc,
				Image:       image,
				Link:        link,
				Source:      "alfavita.gr",
				Title:       title,
			})
		}
	})

	// New Tag: Nasa
	docNasa.Find("#page > main > section.main__content > div.list-promo > article:nth-child(1) > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			if !testURLReachable(link) {
				link = baseURL + link
			}

			// TItle:
			titleQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > div.list-promo__title")
			docNasa.Find(titleQuery).Each(func(j int, w *goquery.Selection) {
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				title = strings.TrimSpace(title)
			})

			// Image:
			imageQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > figure > img")
			docNasa.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
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
			descQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > div.list-promo__summary")
			docNasa.Find(descQuery).Each(func(p int, u *goquery.Selection) {
				desc = u.Text()
				desc = strings.TrimSpace(desc)
			})

			if !thisNewExistsInAnotherTag(link) {
				NewsDBalfavita = append(NewsDBalfavita, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Source:      "alfavita.gr",
					Title:       title,
				})
			}
		}
	})

	// New Tag: selini
	// New Tag: Nasa
	docMoon.Find("#page > main > section.main__content > div.list-promo > article:nth-child(1) > a").Each(func(i int, s *goquery.Selection) {
		link, ok = s.Attr("href")
		if ok {
			if !testURLReachable(link) {
				link = baseURL + link
			}

			// TItle:
			titleQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > div.list-promo__title")
			docMoon.Find(titleQuery).Each(func(j int, w *goquery.Selection) {
				title = s.Text()
				re := regexp.MustCompile(`\r?\n`)
				title = re.ReplaceAllString(title, " ")
				title = strings.TrimSpace(title)
			})

			// Image:
			imageQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > figure > img")
			docMoon.Find(imageQuery).Each(func(j int, w *goquery.Selection) {
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
			descQuery := fmt.Sprintf("#page > main > section.main__content > div.list-promo > article:nth-child(1) > div.list-promo__summary")
			docMoon.Find(descQuery).Each(func(p int, u *goquery.Selection) {
				desc = u.Text()
				desc = strings.TrimSpace(desc)
			})

			if !thisNewExistsInAnotherTag(link) {
				NewsDBalfavita = append(NewsDBalfavita, News{
					Description: desc,
					Image:       image,
					Link:        link,
					Source:      "alfavita.gr",
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
