/*
	Converts a csv file to json
	The first line is treated as a header and the lines are returned as objects in a list

	usage:
		csv-to-json csv-file.csv json-file.json
*/
package main

import (
	"fmt"
	"github.com/marksmithson/csv-to-json/internal/pkg/generators"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	csvFileName := os.Args[1]
	jsonFileName := os.Args[2]

	if csvFileName == "" || jsonFileName == "" {
		printUsage()
		os.Exit(1)
	}

	csvFile, _ := os.Open(csvFileName)

	json, error := generators.CSVToJSON(csvFile)

	if error != nil {
		log.Fatal(error)
	}

	_ = ioutil.WriteFile(jsonFileName, json, 0774)
}

func printUsage() {
	fmt.Println("Usage: csv-to-json csv-file.csv json-file.json")
}
