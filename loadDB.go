package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadDatabase() {

	logActions("Loading database")

	jsonData, err := ioutil.ReadFile("wordListDatabase.json")
	noDatabaseFile := errCheck(err, os.IsNotExist)
	if noDatabaseFile {
		logActions("No database found, proceeding without Database")
		return
	}

	err = json.Unmarshal(jsonData, &mainDatabase)
	errCheck(err, isNoErr)
	logActions("Sucessfully Loaded Database")
}

func saveDatabase() {
	logActions("Saving Database")
	jsonData, err := json.Marshal(mainDatabase)
	errCheck(err, isNoErr)

	err = ioutil.WriteFile("wordListDatabase.json", jsonData, 0664)
	errCheck(err, isNoErr)
	logActions("Successfully saved Database")
}
