package astrobot

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// readDBFile reads the dbFile and returns the JSON
func readDBFile(dbFile string) []byte {
	fileJSON, err := ioutil.ReadFile(dbFile) // Read the file
	if err != nil {
		log.Fatalf("Could not read the file.\nError: %s\n", err)
	}
	return fileJSON
}

// JSONtoOldDB reads the fileDB and writes the information to OldDB
func JSONtoOldDB(dbFile string) {
	log.Println("Saving previous newsposts into OldDB ...")
	fileJSON := readDBFile(dbFile)
	json.Unmarshal(fileJSON, &OldDB)
}
