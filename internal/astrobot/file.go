package astrobot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var postFilePath string = "/Users/drpaneas/github/starlordgr/content/english/post/"

// GetFilename takes a URL as input and a name. It returns the filename with the extension.
func GetFilename(downloadLink, title string) string {
	extension := filepath.Ext(downloadLink)
	title = strings.ReplaceAll(title, " ", "_") // Replace space with underscore
	title = strings.ReplaceAll(title, ",", "_") // Replace comma with underscore
	title = strings.ReplaceAll(title, "&", "_") // Replace & with underscore
	title = strings.ReplaceAll(title, "!", "_") // Replace ! with underscore
	title = strings.ReplaceAll(title, "-", "_") // Replace - with underscore
	title = strings.ReplaceAll(title, ":", "_") // Replace : with underscore
	title = strings.ReplaceAll(title, "?", "_") // Replace ? with underscore
	title = strings.ReplaceAll(title, ";", "_") // Replace ; with underscore
	title = strings.ReplaceAll(title, "\"", "") // Replace ; with underscore
	title = strings.ReplaceAll(title, "#", "")  // Replace ; with underscore

	filename := fmt.Sprintf("%s%s", title, extension)
	return filename
}

func constructFilenamePost(title string) string {
	title = strings.ReplaceAll(title, " ", "_") // Replace space with underscore
	title = strings.ReplaceAll(title, ",", "_") // Replace comma with underscore
	title = strings.ReplaceAll(title, "&", "_") // Replace & with underscore
	title = strings.ReplaceAll(title, "!", "_") // Replace ! with underscore
	title = strings.ReplaceAll(title, "-", "_") // Replace - with underscore
	title = strings.ReplaceAll(title, ":", "_") // Replace : with underscore
	title = strings.ReplaceAll(title, "?", "_") // Replace ? with underscore
	title = strings.ReplaceAll(title, ";", "_") // Replace ; with underscore
	title = strings.ReplaceAll(title, "\"", "") // Replace ; with underscore
	title = strings.ReplaceAll(title, "#", "")  // Replace ; with underscore
	return title + ".md"
}

func writeFile(fullfilepath, content string) {
	fmt.Println(fullfilepath)
	f, err := os.Create(fullfilepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func constructPostFilePath(filename string) string {
	return postFilePath + filename
}

func fixTitle(title string) string {
	if strings.Contains(title, "\"") {
		title = strings.ReplaceAll(title, "\"", "")
	}
	if strings.Contains(title, "«") {
		title = strings.ReplaceAll(title, "«", "")
	}
	if strings.Contains(title, "»") {
		title = strings.ReplaceAll(title, "»", "")
	}
	return title
}

// AddFile creates a new post content from the NewsDB and then it saves the file into the disk.
func AddFile(title, image, source, description, link string) {
	currentTime := time.Now()
	date := fmt.Sprintf("%s", currentTime.Format("2006-01-02T15:04:05-07:00")) // ISO 8601 (RFC 3339)
	title = fixTitle(title)
	content := fmt.Sprintf("---\ntitle: \"%s\"\ndate: %s\nimages:\n  - \"images/post/%s\"\nauthor: \"AstroBot\"\ncategories: [\"Ειδήσεις\"]\ntags: [\"%s\"]\ndraft: false\n---\n\n%s\n\nΔιαβάστε περισσότερα: %s\n", title, date, image, source, description, link)
	filename := constructFilenamePost(title)
	filepath := constructPostFilePath(filename)
	writeFile(filepath, content)
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// GetFileDBPath returns the DB File exact path given its filename
func GetFileDBPath(filename string) string {
	// Find the homedir and create the file
	dbFileNameBefore := "news.before"
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("couldn't find the $HOME directory\nError: %s", err)

	}
	dbFile := home + "/" + dbFileNameBefore
	return dbFile
}

// SaveDBFile saves the NewsDB into the hard disk (that is dbFile location)
func SaveDBFile(dbFile string) {
	fileJSON, err := json.Marshal(NewsDB)
	if err != nil {
		log.Fatal("Couldn't encode to JSON")
	}
	err = ioutil.WriteFile(dbFile, fileJSON, 0644)
	if err != nil {
		log.Fatalf("Couldn't update the db file %s\nError: %s", dbFile, err)
	}
}
