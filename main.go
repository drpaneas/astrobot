package main

import (
	"log"
	"os"

	"github.com/drpaneas/astrobot/internal/astrobot"
)

const (
	filenameDB string = "news.before"
)

func main() {

	dbFile := astrobot.GetFileDBPath(filenameDB)

	// If the file exists parse it (should be in JSON)
	if astrobot.FileExists(dbFile) {

		// Save its contents to OldDB
		astrobot.JSONtoOldDB(dbFile)

		// Get current news in NewsDB
		astrobot.GetCurrentNews()

		// Check if there are any newer news compared to the previous time (run)
		log.Println("Diffing the two databases ...")
		hasNews := astrobot.HasAnyDifference()
		if hasNews {
			astrobot.SaveDBFile(dbFile) // Replace the dbFile with a new one
			astrobot.CreateNewPosts()   // Write the newposts
		} else {
			log.Println("There 0 news since last time we checked")
			os.Exit(0)
		}
	} else {
		// If the fileDB doesn't exist, fetch the news into NewsDB
		astrobot.GetCurrentNews()

		// Encode NewsDB contents in JSON format and save them into dbFile
		astrobot.SaveDBFile(dbFile)
	}
}
