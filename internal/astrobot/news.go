package astrobot

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/drpaneas/astrobot/pkg/cnn"
	"github.com/drpaneas/astrobot/pkg/earthsky"
	"github.com/drpaneas/astrobot/pkg/ecozen"
	"github.com/drpaneas/astrobot/pkg/ert"
	"github.com/drpaneas/astrobot/pkg/in"
	"github.com/drpaneas/astrobot/pkg/maxmag"
	"github.com/drpaneas/astrobot/pkg/naftermporiki"
	"github.com/drpaneas/astrobot/pkg/news247"
	"github.com/drpaneas/astrobot/pkg/newsgr"
	"github.com/drpaneas/astrobot/pkg/protothema"
	"github.com/drpaneas/astrobot/pkg/space"
	"github.com/drpaneas/astrobot/pkg/sputniknews"
	"github.com/drpaneas/astrobot/pkg/tanea"
	"github.com/drpaneas/astrobot/pkg/translate"
	"github.com/drpaneas/astrobot/pkg/unboxholics"
	"github.com/drpaneas/astrobot/pkg/universetoday"
)

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	GreekTitle  string `json:"greektitle"`
	GreekDesc   string `json:"greekdesc"`
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
	log.Println("translate.NewsSpacego()")
	translate.NewsSpacego()
	for _, v := range space.NewsDBSpace {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			GreekTitle:  v.GreekTitle,
			Description: v.Description,
			GreekDesc:   v.GreekDesc,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// EarthSky
	log.Println("earthsky.GetNews()")
	earthsky.GetNews()
	log.Println("translate.NewsEarthSkygo()")
	translate.NewsEarthSkygo()
	for _, v := range earthsky.NewsDBEarthSky {
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			GreekTitle:  v.GreekTitle,
			Description: v.Description,
			GreekDesc:   v.GreekDesc,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// UniverseToday
	log.Println("universetoday.GetNews()")
	universetoday.GetNews()
	log.Println("translate.NewsDBgo()")
	translate.NewsDBgo()
	for _, v := range universetoday.NewsDBUniverseToday {
		if strings.Contains(v.Title, "Hangout") {
			continue
		}
		if v.Title == "" {
			continue
		}
		NewsDB = append(NewsDB, News{
			Title:       v.Title,
			GreekTitle:  v.GreekTitle,
			Description: v.Description,
			GreekDesc:   v.GreekDesc,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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
			GreekTitle:  v.Title,
			Description: v.Description,
			GreekDesc:   v.Description,
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

// HasAnyDifference returns true if NewsDB has newsposts which are not part of the OldDB
// and puts them into DiffDB
func HasAnyDifference() bool {
	thereIsDiff := false
	for _, v := range NewsDB {
		if IsTitleExistsInOldDB(v.Title) {
			continue
		} else {
			fmt.Printf("Title: %s\nDescription: %s\nLink: %s\nImage: %s\n\n", v.Title, v.Description, v.Link, v.Image)
			thereIsDiff = true // flag we gonna have new stuff today
			// Save the current new (which was posted before) into the DiffDB
			DiffDB = append(DiffDB, News{
				Title:       v.Title,
				GreekTitle:  v.GreekTitle,
				Description: v.Description,
				GreekDesc:   v.GreekDesc,
				Link:        v.Link,
				Image:       v.Image,
				Source:      v.Source,
			})
		}
	}
	return thereIsDiff
}

// CreateNewPosts writes news taken from DiffDB
func CreateNewPosts() {
	for _, v := range DiffDB {
		filename := GetFilename(v.Image, v.Title)

		t := time.Now()
		timer := t.Format("20060102150405")
		branch := fmt.Sprintf("news_%s", timer)
		ChangeBranch(branch)

		DownloadImage(v.Image, v.Title)
		AddFile(v.GreekTitle, filename, v.Source, v.GreekDesc, v.Link)

		GitAdd()
		GitCommit()
		GitPush(branch)
		CheckoutMaster()
		time.Sleep(1 * time.Second)
	}
}
