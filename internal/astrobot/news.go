package astrobot

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/drpaneas/astrobot/pkg/alfavita"
	"github.com/drpaneas/astrobot/pkg/businessdaily"
	"github.com/drpaneas/astrobot/pkg/cnn"
	"github.com/drpaneas/astrobot/pkg/earthsky"
	"github.com/drpaneas/astrobot/pkg/ecozen"
	"github.com/drpaneas/astrobot/pkg/egno"
	"github.com/drpaneas/astrobot/pkg/ert"
	"github.com/drpaneas/astrobot/pkg/esquire"
	"github.com/drpaneas/astrobot/pkg/ethnos"
	"github.com/drpaneas/astrobot/pkg/gazzetta"
	"github.com/drpaneas/astrobot/pkg/huffpost"
	"github.com/drpaneas/astrobot/pkg/iefimerida"
	"github.com/drpaneas/astrobot/pkg/in"
	"github.com/drpaneas/astrobot/pkg/maxmag"
	"github.com/drpaneas/astrobot/pkg/naftermporiki"
	"github.com/drpaneas/astrobot/pkg/news247"
	"github.com/drpaneas/astrobot/pkg/newsbomb"
	"github.com/drpaneas/astrobot/pkg/newsgr"
	"github.com/drpaneas/astrobot/pkg/physicsgg"
	"github.com/drpaneas/astrobot/pkg/pontosnews"
	"github.com/drpaneas/astrobot/pkg/protothema"
	"github.com/drpaneas/astrobot/pkg/skai"
	"github.com/drpaneas/astrobot/pkg/space"
	"github.com/drpaneas/astrobot/pkg/sputniknews"
	"github.com/drpaneas/astrobot/pkg/tanea"
	"github.com/drpaneas/astrobot/pkg/thermis"
	"github.com/drpaneas/astrobot/pkg/tovima"
	"github.com/drpaneas/astrobot/pkg/unboxholics"
	"github.com/drpaneas/astrobot/pkg/universetoday"
	"github.com/ecnepsnai/discord"
)

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// NewsDB is the list of news
var NewsDB []News

// OldDB is the list of old news
var OldDB []News

// DiffDB is the diff between the two
var DiffDB []News

// GetCurrentNews fetches current news into the NewsDB database
func GetCurrentNews() {
	log.Println("Saving new newsposts in NewsDB ...")

	// Space.com
	log.Println("space.GetNews()")
	space.GetNews()
	for _, v := range space.NewsDBSpace {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// EarthSky
	log.Println("earthsky.GetNews()")
	earthsky.GetNews()
	for _, v := range earthsky.NewsDBEarthSky {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// UniverseToday
	log.Println("universetoday.GetNews()")
	universetoday.GetNews()
	// log.Println("translate.NewsDBgo()")
	// translate.NewsDBgo()
	for _, v := range universetoday.NewsDBUniverseToday {
		if strings.Contains(v.Title, "Hangout") {
			continue
		}
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// Unoboxholics
	log.Println("unboxholics.GetNews()")
	unboxholics.GetNews()
	for _, v := range unboxholics.NewsDBUnboxholics {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// Naftemporiki
	log.Println("naftermporiki.GetNews()")
	naftermporiki.GetNews()
	for _, v := range naftermporiki.NewsDBNaftermporiki {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// news247
	news247.GetNews()
	for _, v := range news247.NewsDBNews247 {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// TaNea
	log.Println("tanea.GetNews()")
	tanea.GetNews()
	for _, v := range tanea.NewsDBTaNea {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// protothema
	protothema.GetNews()
	for _, v := range protothema.NewsDBProtoThema {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// sputniknews
	log.Println("sputniknews.GetNews()")
	sputniknews.GetNews()
	for _, v := range sputniknews.NewsDBsputnik {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// CNN.gr
	log.Println("cnn.GetNews()")
	cnn.GetNews()
	for _, v := range cnn.NewsDBcnn {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// ert.gr
	log.Println("ert.GetNews()")
	ert.GetNews()
	for _, v := range ert.NewsDBErt {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// in.gr
	log.Println("in.GetNews()")
	in.GetNews()
	for _, v := range in.NewsDBin {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// news.gr
	log.Println("newsgr.GetNews()")
	newsgr.GetNews()
	for _, v := range newsgr.NewsDBnewsgr {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// maxmag.gr
	log.Println("maxmag.GetNews()")
	maxmag.GetNews()
	for _, v := range maxmag.NewsDBmaxmag {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// ecozen.gr
	log.Println("ecozen.GetNews()")
	ecozen.GetNews()
	for _, v := range ecozen.NewsDBecozen {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// gazzetta.gr
	log.Println("gazzetta.GetNews()")
	gazzetta.GetNews()
	for _, v := range gazzetta.NewsDBgazzetta {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// huffingtonpost.gr
	log.Println("huffpost.GetNews()")
	huffpost.GetNews()
	for _, v := range huffpost.NewsDBhuffingtonpost {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// esquire.com.gr
	log.Println("esquire.GetNews()")
	esquire.GetNews()
	for _, v := range esquire.NewsDBesquire {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// tovima.gr
	log.Println("tovima.GetNews()")
	tovima.GetNews()
	for _, v := range tovima.NewsDBtovima {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}
	// iefimerida.gr
	log.Println("iefimerida.GetNews()")
	iefimerida.GetNews()
	for _, v := range iefimerida.NewsDBiefimerida {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// skai.gr
	log.Println("skai.GetNews()")
	skai.GetNews()
	for _, v := range skai.NewsDBskai {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// physicsgg.me
	log.Println("physicsgg.GetNews()")
	physicsgg.GetNews()
	for _, v := range physicsgg.NewsDBphysicsgg {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// ethnos.gr
	log.Println("ethnos.GetNews()")
	ethnos.GetNews()
	for _, v := range ethnos.NewsDBethnos {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// newbomb.gr
	log.Println("newsbomb.GetNews()")
	newsbomb.GetNews()
	for _, v := range newsbomb.NewsDBnewsbomb {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// businessdaily.gr
	log.Println("businessdaily.GetNews()")
	businessdaily.GetNews()
	for _, v := range businessdaily.NewsDBbusinessdaily {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// pontos-news.gr
	log.Println("pontosnews.GetNews()")
	pontosnews.GetNews()
	for _, v := range pontosnews.NewsDBpontos {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// thermisnews.gr
	log.Println("thermis.GetNews()")
	thermis.GetNews()
	for _, v := range thermis.NewsDBthermis {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// egno.gr
	log.Println("egno.GetNews()")
	egno.GetNews()
	for _, v := range egno.NewsDBegno {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// alfavita.gr
	log.Println("alfavita.GetNews()")
	alfavita.GetNews()
	for _, v := range alfavita.NewsDBalfavita {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}
}

// IsTitleExistsInOldDB returns true if title exists in OldDB
func IsTitleExistsInOldDB(title string) bool {
	for _, v := range OldDB {
		if title == v.Title {
			return true
		}
	}
	return false
}

// IsLinkExistsInOldDB returns true if link exists in OldDB
func IsLinkExistsInOldDB(link string) bool {
	for _, v := range OldDB {
		if link == v.Link {
			return true
		}
	}
	return false
}

func fixDesc(desc string) string {
	var re = regexp.MustCompile(`(^|[^&])&#([0-9]){2};([a-z])(|$)`)
	s := re.ReplaceAllString(desc, `$1`)
	return s
}

func fixImageLink(link string) string {
	if strings.Contains(link, "_.") {
		link = strings.Replace(link, "_.", ".", -1)
	}
	return link
}

// HasAnyDifference returns true if NewsDB has newsposts which are not part of the OldDB
// and puts them into DiffDB
func HasAnyDifference() bool {
	thereIsDiff := false
	for _, v := range NewsDB {
		// if IsTitleExistsInOldDB(v.Title) {
		if IsLinkExistsInOldDB(v.Link) {
			continue
		} else {
			fmt.Printf("Title: %s\nDescription: %s\nLink: %s\nImage: %s\n\n", v.Title, v.Description, v.Link, v.Image)
			thereIsDiff = true // flag we gonna have new stuff today
			// Save the current new (which was posted before) into the DiffDB
			DiffDB = append(DiffDB, News{
				Title:       v.Title,
				Description: fixDesc(v.Description),
				Link:        v.Link,
				Image:       fixImageLink(v.Image),
				Source:      v.Source,
			})
		}
	}
	return thereIsDiff
}

func isGreek(source string) bool {
	sources := [...]string{
		"huffingtonpost.gr",
		"ecozen.gr",
		"news247.gr",
		"in.gr",
		"tovima.gr",
		"unboxholics.com",
		"ert.gr",
		"esquire.com.gr",
		"naftemporiki.gr",
		"maxmag.gr",
		"cnn.gr",
		"tanea.gr",
		"protothema.gr",
		"gazzetta.gr",
		"sputniknews.gr",
		"news.gr",
		"iefimerida.gr",
		"skai.gr",
		"physicsgg.me",
		"ethnos.gr",
		"newsbomb.gr",
		"businessdaily.gr",
		"pontos-news.gr",
		"thermisnews.gr",
		"egno.gr",
		"alfavita.gr",
	}
	for _, v := range sources {
		if source == v {
			log.Printf("%s is a Greek site. Push directly to master branch.", v)
			return true
		}
	}
	return false
}

func postDiscord(webhook, link, title, desc, imageLink string) error {
	var pic discord.Image
	pic.URL = imageLink
	discord.WebhookURL = webhook
	err := discord.Post(discord.PostOptions{
		// Content: text,
		Embeds: []discord.Embed{
			{
				Color:       16777215,
				URL:         link,
				Title:       title,
				Description: desc,
				Thumbnail:   &pic,
			},
		},
	})
	return err
}

// CreateNewPosts writes news taken from DiffDB
func CreateNewPosts(webhook string) {
	for _, v := range DiffDB {
		IsItUpToDate()
		CheckoutMaster()
		if DownloadImage(v.Image) != nil {
			log.Println("FAILURE !!!!!!!!!!!!!!")
			log.Printf("Problem is found at %v\n\n\n", v)
			continue
		}
		AddFile(v.Title, imageFilename(v.Image), v.Source, v.Description, v.Link, webhook, v.Image)
		if BuildFails() {
			log.Println("FAILURE !!!!!!!!!!!!!!")
			log.Printf("Problem is found at %v\n\n\n", v)
			continue
		}
		GitAdd()
		GitCommit()
		GitPush("master")
		wait := 5 * time.Second
		log.Printf("\n -- Waiting %v seconds -- \n", wait)
		time.Sleep(wait)

		fmt.Printf("Done\n")
	}
}
