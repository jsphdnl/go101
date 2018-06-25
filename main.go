package main

import (
	"fmt"
	"exercises/quiz1/cli"
	"exercises/quiz1/csvparser"
	"exercises/quiz1/core"
	"os"
	"log"
	"encoding/json"
)

const (
	DEFAULT_FILE_NAME = "problems.csv"
)

// Entry point
// Usage :  optional --file-name <filename>
// 	    If not specified will default to "problems.csv" in
//          the current location
func main() {
	//Handle interrupts
	//Check args
	args := os.Args
	if !(len(args) == 1 || len(args) == 2) {
		cli.Usage()
	}
	fileName := DEFAULT_FILE_NAME
	if len(args) == 2 {
		fileName = os.Args[1]
	}

	records, err := csvparser.Parse(fileName)
	if err != nil {
		log.Fatal(err)
	}
	score := core.Startquiz(records)
	sbr, _ := json.Marshal(&score)
	fmt.Println(string(sbr))
}
