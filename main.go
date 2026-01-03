package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv`)
	}
}
func main() {

	flag.Parse()

	if flag.NArg() != 3 {
		flag.Usage()
		os.Exit(1)
	}

	input := flag.Arg(0)
	//output := flag.Arg(1)//
	airportLookup := flag.Arg(2)

	checkFileExist(input, airportLookup)
	//CSV file opening and validation
	fileOutput, err := openFiles(airportLookup)
	if err != nil {
		fmt.Fprint(os.Stderr, "Airport lookup not found")
		os.Exit(1)
	}

	if !validCSV(fileOutput) {
		fmt.Fprint(os.Stderr, "Airport lookup malformed")
		os.Exit(1)
	}

	//Input file reading
	fileInput, err := inputTextRead(input)
	if err != nil {
		fmt.Fprint(os.Stderr, "Input not found")
		os.Exit(1)
	}

	//trim or clean kora input data store kora

	var trimmedLines []string // new cleaned lines store korar jonno
	lastLineEmpty := false
	//fileInput onekgulo sentencer er slice or talika r line hoilo eikhane ekta sentence (string)
	for _, line := range fileInput {

		trimmed := strings.TrimSpace(line) //whitespace gulo remove korbe

		//advance method use kora lagbe emon kono ulta palta jinish khoj korar jonno-normal kaj korar jonno strings package use korai enough
		replaceChar := trimmed
		replaceChar = strings.ReplaceAll(replaceChar, "\r", "\n")
		replaceChar = strings.ReplaceAll(replaceChar, "\v", "\n")
		replaceChar = strings.ReplaceAll(replaceChar, "\f", "\n")

		for _, singleLine := range strings.Split(replaceChar, "\n") {

			//multiple empty line check
			if singleLine == "" && lastLineEmpty {
				continue
			}
			if singleLine == "" {
				lastLineEmpty = true
			} else {
				lastLineEmpty = false
			}
			//empty space check
			trimmedLines = append(trimmedLines, singleLine)
		}
	}
}

func checkFileExist(input string, airportLookup string) {
	_, err := os.Stat(input)
	_, err1 := os.Stat(airportLookup)

	if os.IsNotExist(err) {
		fmt.Fprint(os.Stderr, "Input not found\n")
		os.Exit(1)
	}
	if os.IsNotExist(err1) {
		fmt.Fprint(os.Stderr, "Airport Lookup not found\n")
		os.Exit(1)
	}
}

func openFiles(path string) ([][]string, error) {
	//open the file
	fileCsv, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fileCsv.Close() //close the file after function return-protecting from memory leak

	//read csv file
	reader := csv.NewReader(fileCsv)
	return reader.ReadAll() //inside csv everything is just text no data types
}

var expectedHeader = []string{
	"name",
	"iso_country",
	"municipality",
	"icao_code",
	"iata_code",
	"coordinates",
}

// csv validation
func validCSV(reader [][]string) bool {
	//header validation
	header := reader[0]
	if len(header) != 6 {

		return false
	}
	for i, v := range header {
		if v != expectedHeader[i] {

			return false
		}

	}

	//row validation
	for i := 1; i < len(reader); i++ {
		row := reader[i]
		if len(row) != 6 {

			return false
		}
		for _, column := range row {
			if len(strings.TrimSpace(column)) == 0 {

				return false
			}

		}

	}
	return true
}

//input text reading

func inputTextRead(path string) ([]string, error) {

	fileTxt, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fileTxt.Close()

	scanner := bufio.NewScanner(fileTxt)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}
