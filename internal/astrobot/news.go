package astrobot

import (
	"fmt"
	"github.com/drpaneas/astrobot/pkg/alfavita"
	"github.com/drpaneas/astrobot/pkg/astronio"
	"github.com/drpaneas/astrobot/pkg/businessdaily"
	"github.com/drpaneas/astrobot/pkg/cnn"
	"github.com/drpaneas/astrobot/pkg/earthsky"
	"github.com/drpaneas/astrobot/pkg/ecozen"
	"github.com/drpaneas/astrobot/pkg/egno"
	"github.com/drpaneas/astrobot/pkg/esquire"
	"github.com/drpaneas/astrobot/pkg/ethnos"
	"github.com/drpaneas/dudenetes/pkg/run"

	//	"github.com/drpaneas/astrobot/pkg/gazzetta"
	"github.com/drpaneas/astrobot/pkg/iefimerida"
	"github.com/drpaneas/astrobot/pkg/in"
	"github.com/drpaneas/astrobot/pkg/maxmag"
	"github.com/drpaneas/astrobot/pkg/naftermporiki"
	"github.com/drpaneas/astrobot/pkg/nasapicofday"
	"github.com/drpaneas/astrobot/pkg/news247"
	"github.com/drpaneas/astrobot/pkg/newsbomb"
	"github.com/drpaneas/astrobot/pkg/newsgr"
	"github.com/drpaneas/astrobot/pkg/physicsgg"
	//"github.com/drpaneas/astrobot/pkg/pontosnews"
	"github.com/drpaneas/astrobot/pkg/protothema"
	"github.com/drpaneas/astrobot/pkg/skai"
	"github.com/drpaneas/astrobot/pkg/space"
	"github.com/drpaneas/astrobot/pkg/sputniknews"
	"github.com/drpaneas/astrobot/pkg/tanea"
	"github.com/drpaneas/astrobot/pkg/thermis"
	"github.com/drpaneas/astrobot/pkg/tovima"
	//"github.com/drpaneas/astrobot/pkg/unboxholics" disabled
	"github.com/drpaneas/astrobot/pkg/universetoday"
	"github.com/ecnepsnai/discord"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// News represent an news article
type News struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Source      string `json:"source"`
}

// NewDB is the list of news
var NewDB []News

// OldDB is the list of old news
var OldDB []News

// DiffDB is the diff between the two
var DiffDB []News

// TestedDB contains the good news (after test) taken from DiffDB
var TestedDB []News

// GetCurrentNews fetches current news into the NewDB database
func GetCurrentNews() {
	log.Println("Saving new newsposts in NewDB ...")

	// Space.com
	log.Println("space.GetNews()")
	space.GetNews()
	for _, v := range space.NewsDBSpace {
		if v.Title == "" {
			continue
		}
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// Nasa Picture of the day
	log.Println("nasapicofday.GetNews()")
	nasapicofday.GetNews()
	for _, v := range nasapicofday.NewsDBNasaImage {
		if v.Title == "" {
			continue
		}
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// Unoboxholics
	//log.Println("unboxholics.GetNews()")
	//unboxholics.GetNews()
	//for _, v := range unboxholics.NewsDBUnboxholics {
	//	if v.Title == "" {
	//		continue
	//	}
	//	NewDB = append(NewDB, News{
	//		Title:       v.Title,
	//		Description: v.Description,
	//		Link:        v.Link,
	//		Image:       v.Image,
	//		Source:      v.Source,
	//	})
	//}

	// Astronio.gr
	log.Println("astronio.GetNews()")
	astronio.GetNews()
	for _, v := range astronio.NewsDBastronio {
		if v.Title == "" {
			continue
		}
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// ert.gr
	// log.Println("ert.GetNews()")
	// ert.GetNews()
	// for _, v := range ert.NewsDBErt {
	// 	if v.Title == "" {
	// 		continue
	// 	}
	// 	NewDB = append(NewDB, News{
	// 		Title:       v.Title,
	// 		Description: v.Description,
	// 		Link:        v.Link,
	// 		Image:       v.Image,
	// 		Source:      v.Source,
	// 	})
	// }

	// in.gr
	log.Println("in.GetNews()")
	in.GetNews()
	for _, v := range in.NewsDBin {
		if v.Title == "" {
			continue
		}
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// gazzetta.gr
	//log.Println("gazzetta.GetNews()")
	//gazzetta.GetNews()
	//for _, v := range gazzetta.NewsDBgazzetta {
	//	if v.Title == "" {
	//		continue
	//	}
	//	NewDB = append(NewDB, News{
	//		Title:       v.Title,
	//		Description: v.Description,
	//		Link:        v.Link,
	//		Image:       v.Image,
	//		Source:      v.Source,
	//	})
	//}

	// huffingtonpost.gr
	// log.Println("huffpost.GetNews()")
	// huffpost.GetNews()
	// for _, v := range huffpost.NewsDBhuffingtonpost {
	// 	if v.Title == "" {
	// 		continue
	// 	}
	// 	NewDB = append(NewDB, News{
	// 		Title:       v.Title,
	// 		Description: v.Description,
	// 		Link:        v.Link,
	// 		Image:       v.Image,
	// 		Source:      v.Source,
	// 	})
	// }

	// esquire.com.gr
	log.Println("esquire.GetNews()")
	esquire.GetNews()
	for _, v := range esquire.NewsDBesquire {
		if v.Title == "" {
			continue
		}
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})
	}

	// pontos-news.gr
	// log.Println("pontosnews.GetNews()")
	// pontosnews.GetNews()
	// for _, v := range pontosnews.NewsDBpontos {
	//	if v.Title == "" {
	//		continue
	//	}
	//	NewDB = append(NewDB, News{
	//		Title:       v.Title,
	//		Description: v.Description,
	//		Link:        v.Link,
	//		Image:       v.Image,
	//		Source:      v.Source,
	//	})
	// }

	// thermisnews.gr
	log.Println("thermis.GetNews()")
	thermis.GetNews()
	for _, v := range thermis.NewsDBthermis {
		if v.Title == "" {
			continue
		}
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
		NewDB = append(NewDB, News{
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
	for _ , v := range OldDB {
		// fmt.Printf("OldDB[%d].Link: %s\n",i,v.Link)
		if link == v.Link {
			// fmt.Println("This is old news. We just found it.")
			return true
		}
		// fmt.Println("Try another OldDB element.")
	}
	// fmt.Println("Couldn't find the same link in OldDB. This looks like fresh news!")
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

// HasAnyDifference returns true if NewDB has newsposts which are not part of the OldDB
// and puts them into DiffDB
func HasAnyDifference() bool {
	thereIsDiff := false
	fmt.Printf("OldDB has %d news\nNewDB has %d news\n", len(OldDB), len(NewDB))
	for _, v := range NewDB {
		// fmt.Printf("NewDB[%d].Link: %s\nCheck if this exists in OldDB\n",i, v.Link)
		// if IsTitleExistsInOldDB(v.Title) {
		if IsLinkExistsInOldDB(v.Link) {
			fmt.Printf("Ignoring '%s' -- already exists in OldDB\n", v.Title)
			continue
		} else {
			fmt.Println("----")
			fmt.Println("There is difference between OldDB and NewDB. This is one of the new elements:")
			fmt.Printf("Title: %s\nDescription: %s\nLink: %s\nImage: %s\n\n", v.Title, fixDesc(v.Description), v.Link, fixImageLink(v.Image))
			fmt.Println("----")
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
		// "huffingtonpost.gr", Disable because of GET  request failed: Get "": unsupported protocol scheme ""
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
		"astronio.gr",
	}
	for _, v := range sources {
		if source == v {
			return true
		}
	}
	return false
}

func postDiscord(webhook, link, title, desc, imageLink, source string) error {
	var pic discord.Image
	pic.URL = imageLink
	discord.WebhookURL = webhook
	var err error
	if source != "nasa.gov" {
		err = discord.Post(discord.PostOptions{
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
	} else {
		// Σημαίνει είμαστε στην Φωτο της Ημέρεας
		err = discord.Post(discord.PostOptions{
			// Content: text,
			Embeds: []discord.Embed{
				{
					Color:       16777215,
					URL:         link,
					Title:       "Η φωτογραφία της ημέρας: " + title,
					Description: desc,
					Image:       &pic,
				},
			},
		})
	}

	return err
}

// TestNewPosts tests news taken from DiffDB
func TestNewPosts() {
	for _, v := range DiffDB {
		//IsItUpToDate()
		//fmt.Println("Before checkout")
		//CheckoutMaster()
		//fmt.Println("After checkout")
		// Αν το συγκεκριμένο νέο έχει πρόβλημα με την εικόνα, προσπέρασέ το

		imageFilepath := constructImageFilePath(removeQuote(imageFilename(v.Image)))

		fmt.Println("--- Testing Image ---")
		fmt.Println("Link: " + v.Image)	// e.g.  https://www.universetoday.com/wp-content/uploads/2021/03/jpegPIA24483.width-1600.jpg
		fmt.Println("Save To: " + imageFilepath) // e.g. /home/runner/work/starlordgr/starlordgr/website/static/images/post/jpegPIA24483.width-1600.jpg
		imageAlreadyPreExists := false // e.g. το σιτε χρησιμοποιεί την ίδια εικόνα σε διάφορες ειδήσεις
		if FileExists(imageFilepath) {
			fmt.Printf("The image %v already exists! Do not download it again ...\n",imageFilepath)
			imageAlreadyPreExists = true
		} else {
			if DownloadImage(v.Image, imageFilepath) != nil {
				log.Println("FAILURE with Downloading the Image: ", v.Image)
				log.Printf("Problem is found at %v\n", v)
				log.Printf("Skipping this news %v\n\n\n", v.Link)
				// Δεν ήταν καλή η φωτογραφία της είδησης. Διέγραψε το αρχείο
				os.Remove(imageFilepath)
				continue
			}
		}

		// Πρόσθεσε το νέο ειδησιογραφικό αρχείο και τέσταρέ το
		title := fixTitle(v.Title)
		filename := constructFilenamePost(title)
		filepath := constructEnglishPostFilePath(filename)
		if isGreek(v.Source) {
			filepath = constructGreekPostFilePath(filename)
		}
		fmt.Println("--- Testing Adding File ---")
		fmt.Println("File: " + filepath)
		if FileExists(filepath) {
			fmt.Println("The file already exists! Skipping ...")
			continue
		}
		AddFile(v.Title, imageFilename(v.Image), v.Source, v.Description, v.Link, v.Image, filepath)

		fmt.Println("--- Testing Building File ---")
		if BuildFails() {
			log.Println("FAILURE during Testing building file !!!!!!!!!!!!!!")
			log.Println("meaning --> 'hugo --gc --themesDir themes' command has failed")
			log.Printf("Problem is found at %v\n", v)
			cmd := fmt.Sprintf("cat %v", filepath)
			log.Println("Running: ", cmd)
			timeout := 60
			output, _ := run.SlowCmdDir(cmd, timeout, directory)
			log.Println("----------------------------------------")
			log.Println(output)
			log.Println("----------------------------------------")
			log.Printf("Skipping this specific new article. Delete the file and the image %v\n\n\n", v.Link)
			// Δεν ήταν καλή η είδηση. Διέγραψε το αρχείο και την εικόνα
			os.Remove(filepath)
			if !imageAlreadyPreExists {
				os.Remove(imageFilepath)	// ειδική περίπτωση όπου η φωτο προυπάρχει (για άλλα άρθρα) οπότε μην την σβήνεις
			}
			continue
		}

		// Αφαιρεσε το άρθρο, όπως και να χει γιατι αυτο ήταν απλά ένα τεστ
		if !imageAlreadyPreExists {
			os.Remove(imageFilepath)	// ειδική περίπτωση όπου η φωτο προυπάρχει (για άλλα άρθρα) οπότε μην την σβήνεις
		}
		os.Remove(filepath)

		// Η είδηση πέρασε όλα τα τεστ, θεωρείται πλέον αξιόπιστη για push
		TestedDB = append(TestedDB, News{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       v.Image,
			Source:      v.Source,
		})

		//GitAdd()
		//GitCommit()
		//GitPush("master")
		// wait := 5 * time.Second
		// log.Printf("\n -- Waiting %v seconds -- \n", wait)
		// time.Sleep(wait)

		fmt.Printf("Testing is Done\n")
	}
}

// CreateNewPosts writes news taken from TestedDB
func CreateNewPosts() {
	for i, v := range TestedDB {
		// Αν το συγκεκριμένο νέο έχει πρόβλημα με την εικόνα, προσπέρασέ το
		imageFilepath := constructImageFilePath(removeQuote(imageFilename(v.Image)))
		imageAlreadyPreExists := false // e.g. το σιτε χρησιμοποιεί την ίδια εικόνα σε διάφορες ειδήσεις
		if FileExists(imageFilepath) {
			fmt.Println("The image already exists! Do not download it again ...")
			imageAlreadyPreExists = true
		} else {
			if DownloadImage(v.Image, imageFilepath) != nil {
				log.Println("FAILURE with Downloading the Image: ", v.Image)
				log.Printf("Problem is found at %v\n", v)
				log.Printf("Skipping this news %v\n\n\n", v.Link)
				// Δεν ήταν καλή η φωτογραφία της είδησης. Διέγραψε το αρχείο
				os.Remove(imageFilepath)
				continue
			}
		}

		// Πρόσθεσε το νέο ειδησιογραφικό αρχείο και τέσταρέ το
		title := fixTitle(v.Title)
		filename := constructFilenamePost(title)
		filepath := constructEnglishPostFilePath(filename)
		if isGreek(v.Source) {
			filepath = constructGreekPostFilePath(filename)
		}
		AddFile(v.Title, imageFilename(v.Image), v.Source, v.Description, v.Link, v.Image, filepath)
		if BuildFails() {
			log.Println("FAILURE during building the file !!!!!!!!!!!!!!")
			log.Printf("Problem is found at %v\n", v)
			cmd := fmt.Sprintf("cat %v", filepath)
			log.Println("Running: ", cmd)
			timeout := 60
			output, _ := run.SlowCmdDir(cmd, timeout, directory)
			log.Println(output)
			log.Printf("Skipping this specific new article %v\n\n\n", v.Link)
			// Δεν ήταν καλή η είδηση. Διέγραψε το αρχείο
			if !imageAlreadyPreExists {
				os.Remove(imageFilepath)	// ειδική περίπτωση όπου η φωτο προυπάρχει (για άλλα άρθρα) οπότε μην την σβήνεις
			}
			os.Remove(filepath)
			continue
		}
		fmt.Printf("Added %d out of %d\n", i+1, len(TestedDB))
	}
	fmt.Printf("Adding all files is Done\n")
}

func PostToDiscord(webhook string) {
	for _, v := range TestedDB {
		time.Sleep(10 * time.Second) // Περίμενε γιατί αν είναι πολλά τα νέα, παίζει να φας ban απο τον server για spam
		if v.Source != "newsbomb.gr" && v.Source != "sputniknews.gr" {
			// Send to Discord Community
			err := postDiscord(webhook, v.Link, v.Title, v.Description, v.Image, v.Source)
			if err != nil {
				fmt.Printf("\n######### Error with Discord #########\n")
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("\nDiscord successfully sent: %s\n", v.Title)
			}
		}
	}
}
