package astrobot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	postEnglishFilePath = GetRepoPath() + "/content/english/post/"
	postGreekFilePath   = GetRepoPath() + "/content/greek/post/"
	imagesFilepath      = GetRepoPath() + "/static/images/post/"
	directory           = GetRepoPath()
)

func GetRepoPath() string {
	repoPath, present := os.LookupEnv("REPO_PATH")
	fmt.Println(repoPath)
	if !present {
		log.Fatal("REPO_PATH env variable does NOT exists")
		os.Exit(1)
	}
	return repoPath
}

// GetFilename takes a URL as input and a name. It returns the filename with the extension.
func GetFilename(downloadLink, title string) string {
	extension := filepath.Ext(downloadLink)
	title = strings.ReplaceAll(title, " ", "")  // Replace space with underscore
	title = strings.ReplaceAll(title, ",", "")  // Replace comma with underscore
	title = strings.ReplaceAll(title, "&", "")  // Replace & with underscore
	title = strings.ReplaceAll(title, "!", "")  // Replace ! with underscore
	title = strings.ReplaceAll(title, "-", "")  // Replace - with underscore
	title = strings.ReplaceAll(title, ":", "")  // Replace : with underscore
	title = strings.ReplaceAll(title, "?", "")  // Replace ? with underscore
	title = strings.ReplaceAll(title, ";", "")  // Replace ; with underscore
	title = strings.ReplaceAll(title, "\"", "") // Replace \ with underscore
	title = strings.ReplaceAll(title, "#", "")  // Replace # with underscore

	filename := fmt.Sprintf("%s%s", title, extension)
	return filename
}

// FixImageFilename fixes the image filaname
// func FixImageFilename(image string) string {
// 	image = strings.ReplaceAll(image, " ", "") // Replace space with underscore
// 	image = strings.ReplaceAll(image, ",", "") // Replace comma with underscore
// 	image = strings.ReplaceAll(image, "&", "") // Replace & with underscore
// 	image = strings.ReplaceAll(image, "!", "") // Replace ! with underscore
// 	//image = strings.ReplaceAll(image, "-", "_") // Replace - with underscore
// 	image = strings.ReplaceAll(image, ":", "")  // Replace : with underscore
// 	image = strings.ReplaceAll(image, "?", "")  // Replace ? with underscore
// 	image = strings.ReplaceAll(image, ";", "")  // Replace ; with underscore
// 	image = strings.ReplaceAll(image, "\"", "") // Replace \ with underscore
// 	image = strings.ReplaceAll(image, "#", "")  // Replace # with underscore
// 	image = strings.ReplaceAll(image, "?", "")  // Replace # with underscore
// 	return image
// }

func constructFilenamePost(title string) string {
	title = strings.ReplaceAll(title, " ", "")  // Replace space with underscore
	title = strings.ReplaceAll(title, ",", "")  // Replace comma with underscore
	title = strings.ReplaceAll(title, "&", "")  // Replace & with underscore
	title = strings.ReplaceAll(title, "!", "")  // Replace ! with underscore
	title = strings.ReplaceAll(title, "-", "")  // Replace - with underscore
	title = strings.ReplaceAll(title, ":", "")  // Replace : with underscore
	title = strings.ReplaceAll(title, "?", "")  // Replace ? with underscore
	title = strings.ReplaceAll(title, ";", "")  // Replace ; with underscore
	title = strings.ReplaceAll(title, "\"", "") // Replace \ with underscore
	title = strings.ReplaceAll(title, "#", "")  // Replace # with underscore
	return title + ".md"
}

func writeFile(fullfilepath, content string) error {
	fmt.Println(fullfilepath)
	dir := path.Dir(fullfilepath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil { // Use MkdirAll to simulate mkdir -p
			log.Panicf("I couldn't create the directory to save the markdown files: %v", err)
		} else {
			fmt.Println("Directory to save markdown files has been created:", dir)
		}
	}
	f, err := os.Create(fullfilepath)
	if err != nil {
		fmt.Println("Error with os.Create:", err)
		return err
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println("Error with f.WriteString:", err)
		f.Close()
		return err
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println("Error with f.Close: ", err)
		return err
	}
	return nil
}

func constructEnglishPostFilePath(filename string) string {
	return postEnglishFilePath + filename
}
func constructGreekPostFilePath(filename string) string {
	return postGreekFilePath + filename
}

func fixTitle(title string) string {
	if strings.Contains(title, "\"") {
		title = strings.ReplaceAll(title, "\"", "")
	}
	if strings.Contains(title, "/") {
		title = strings.ReplaceAll(title, "/", "")
	}
	if strings.Contains(title, "\\") {
		title = strings.ReplaceAll(title, "\\", "")
	}
	if strings.Contains(title, "«") {
		title = strings.ReplaceAll(title, "«", "")
	}
	if strings.Contains(title, "»") {
		title = strings.ReplaceAll(title, "»", "")
	}
	if strings.Contains(title, "#") {
		title = strings.ReplaceAll(title, "#", "")
	}
	if strings.Contains(title, ":") {
		title = strings.ReplaceAll(title, ":", "")
	}
	return title
}

// AddFile creates a new post content from the NewDB and then it saves the file into the disk.
func AddFile(title, image, source, description, link, imageLink, filepath string) error {
	currentTime := time.Now()
	date := fmt.Sprintf("%s", currentTime.Format("2006-01-02T15:04:05-07:00")) // ISO 8601 (RFC 3339)
	category := "News"
	if isGreek(source) {
		category = "Ειδήσεις"
	}
	content := fmt.Sprintf("---\ntitle: \"%s\"\ndate: %s\nimages:\n  - \"images/post/%s\"\nauthor: \"AstroBot\"\ncategories: [\"%s\"]\ntags: [\"%s\"]\ndraft: false\n---\n\n%s\n\nΔιαβάστε περισσότερα: %s\n", title, date, image, category, source, description, link)
	if err := writeFile(filepath, content); err != nil {
		fmt.Println("Failed to writeFile() -- cannot save the markdown file to disk")
		return err
	}
	fmt.Println("Markdown file has been saved to disk!")
	return nil
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// GetFileDBPath returns the DB File exact path given its filename
func GetFileDBPath(filename string) string {
	dbFile := directory + "/" + filename
	if !FileExists(dbFile) {
		log.Fatalf("The file %s cannot be found!\n", dbFile)
	}
	return dbFile
}

// SaveDBFile saves the NewDB into the hard disk (that is dbFile location)
// I use NewDB because I don't want to retest the problematic new articles every time the CI gets triggered
func SaveDBFile(dbFile string) {
	// Append OldDB with the new added stuff
	for _, v := range NewDB {
		OldDB = append(OldDB, v)
	}

	// Remove duplicates if any
	OldDB = uniqueDB(OldDB)

	// The OldDB is now bigger
	fileJSON, err := json.Marshal(OldDB)
	if err != nil {
		log.Fatal("Couldn't encode to JSON")
	}
	err = ioutil.WriteFile(dbFile, fileJSON, 0644)
	if err != nil {
		log.Fatalf("Couldn't update the db file %s\nError: %s", dbFile, err)
	}
}

func uniqueDB(NewSlice []News) []News {
	keys := make(map[News]bool)
	var list []News
	for _, entry := range NewSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
