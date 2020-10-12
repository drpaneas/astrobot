package astrobot

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/drpaneas/astrobot/pkg/earthsky"
	"github.com/drpaneas/astrobot/pkg/space"
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
	space.GetNews()
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
	earthsky.GetNews()
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
	universetoday.GetNews()
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
