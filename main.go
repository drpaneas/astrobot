package main

import (
	"fmt"
	"log"
	"os"

	"github.com/drpaneas/astrobot/internal/astrobot"
)

//func gitGoodState() {
//	astrobot.IsItUpToDate() // Pull Changes (if there are any)
//	if !astrobot.IsItUpToDate() {
//		log.Fatal("Repo is not clean") // There should not be any changes now
//	}
//}

const (
	filenameDB string = "news.before"
)

var (
	webhook = astrobot.GetWebhook()
)

func main() {
	// gitGoodState()

	dbFile := astrobot.GetFileDBPath(filenameDB)

	// If the file exists parse it (should be in JSON)
	if astrobot.FileExists(dbFile) {
		fmt.Printf("Database file found: %s\n", dbFile)

		// Save its contents to OldDB
		astrobot.JSONtoOldDB(dbFile)

		// Get current news in NewDB
		astrobot.GetCurrentNews()

		// Check if there are any newer news compared to the previous time (run)
		log.Println("Diffing the two databases ...")
		hasNews := astrobot.HasAnyDifference()
		if hasNews {
			// Για κάθε entry στης NewDB δοκίμασε να δεις αν κάνει build
			fmt.Println("----- START TESTING ------")
			astrobot.TestNewPosts() // Test the newposts
			fmt.Println("----- FINISHED TESTING ------")
			fmt.Println("----- START ADDING FILES ------")
			// astrobot.CreateNewPosts() // Write the newposts
			fmt.Println("----- FINISH ADDING FILES ------")
			fmt.Println("----- START SENDING DISCORD MESSAGES ------")
			astrobot.PostToDiscord(webhook) // Push messages to the discord room
			fmt.Println("----- FINISH SENDING DISCORD MESSAGES ------")
			fmt.Println("----- START REPLACING THE OLDDB WITH THE TESTEDDB ------")
			// astrobot.SaveDBFile(dbFile)      // Replace the dbFile with a new one
			fmt.Println("----- FINISHED REPLACING THE OLDDB WITH THE TESTEDDB ------")
		} else {
			log.Println("There 0 news since last time we checked")
			os.Exit(0)
		}
	} else {
		fmt.Printf("There is no %s file. We are gathering the news for the first time.\n", dbFile)

		// If the fileDB doesn't exist, fetch the news into NewDB
		astrobot.GetCurrentNews()

		// Encode NewDB contents in JSON format and save them into dbFile
		astrobot.SaveDBFile(dbFile)
	}
}
