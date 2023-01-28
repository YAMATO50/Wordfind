package main

import (
	"encoding/json"
	"os"
)

func loadDatabase() {

	logActions("Loading database")

	jsonData, err := os.ReadFile("wordListDatabase.json")
	noDatabaseFile := errCheck(err, os.IsNotExist)
	if noDatabaseFile {
		logActions("No database found, proceeding with empty database")
		mainDatabase.SameLengthWords = make(map[int]SameLengthWord)
		return
	}

	err = json.Unmarshal(jsonData, &mainDatabase)
	errCheck(err, isNoErr)
	logActions("Sucessfully loaded database")
}

func saveDatabase() {
	logActions("Saving database")
	jsonData, err := json.Marshal(mainDatabase)
	errCheck(err, isNoErr)

	err = os.WriteFile("wordListDatabase.json", jsonData, 0664)
	errCheck(err, isNoErr)
	logActions("Successfully saved database")
}
