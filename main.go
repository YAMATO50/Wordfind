package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var flagBuldDatabase bool
var newBuildDBfile string
var newBuildDBfileExt string
var verbose bool
var mainDatabase Database

func main() {
	getFlags()
	loadDatabase()
	if flagBuldDatabase {
		logActions("Building database")
		timeMeasurement()
		buildDatabase(newBuildDBfile, newBuildDBfileExt)
		elapsed := timeMeasurement()
		logActions(fmt.Sprintf("Database building took %d ns", elapsed))
	}
}

func getFlags() {
	args := os.Args
	for i, arg := range args {
		arg = strings.ToLower(arg)
		switch arg {
		case "-b":
			handleBuildDBFlag(args, i)
		case "-s":
			handleSingleWordFlag(args, i)
		case "-v":
			verbose = true
		default:
			continue
		}
	}
}

func handleBuildDBFlag(args []string, argPos int) {
	flagBuldDatabase = true

	if len(args) <= argPos+1 {
		fmt.Print("Missing filepath after -b flag")
		os.Exit(1)
	}

	newBuildDBfile = args[argPos+1]
	extFields := strings.Split(newBuildDBfile, ".")
	if len(extFields) == 1 {
		fmt.Println("Unsupported file extension")
		os.Exit(1)
	}

	newBuildDBfileExt = strings.ToLower(extFields[len(extFields)-1])
	switch newBuildDBfileExt {
	case "txt":
		break
	default:
		fmt.Println("Unsupported file extension")
		os.Exit(1)
	}
}

func handleSingleWordFlag(args []string, argPos int) {
	flagBuldDatabase = true

	if len(args) <= argPos+1 {
		fmt.Println("Missing at least one word after -s flag")
		os.Exit(1)
	}

	newBuildDBfile = strings.Join(args[argPos+1:], "\n")
	newBuildDBfileExt = "list"
}

func logActions(logString string) {
	if verbose {
		fmt.Println(logString)
	}
}

var firstCall bool = true
var startTime int64

func timeMeasurement() int64 {
	if !verbose {
		return 0
	}

	if !firstCall {
		startTime = time.Now().UnixNano()
		firstCall = false
		return 0
	}

	firstCall = true
	return time.Now().UnixNano() - startTime
}
