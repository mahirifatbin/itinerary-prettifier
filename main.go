package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
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
	// Step 7: Transforming IATA & ICAO codes to Airport Names

	buildLookupMaps(fileOutput)

	//trim or clean kora input data store kora
	//STEP 6 : input text clean kora

	var trimmedLines []string // new cleaned lines store korar jonno
	lastLineEmpty := false
	re := regexp.MustCompile(`[ \t]+`) //multiple space or tab remove korar jonno regexp.loop er baire rakhsi karon bar bar banate hobe na

	//fileInput onekgulo sentencer er slice or talika r line hoilo eikhane ekta sentence (string)
	for _, line := range fileInput {

		trimmed := strings.TrimSpace(line) //samne and pichoner whitespace gulo remove korbe. majher gulo nah

		//advance method use kora lagbe emon kono ulta palta jinish khoj korar jonno-normal kaj korar jonno strings package use korai enough
		replaceChar := trimmed
		replaceChar = strings.ReplaceAll(replaceChar, "\r", "\n")
		replaceChar = strings.ReplaceAll(replaceChar, "\v", "\n")
		replaceChar = strings.ReplaceAll(replaceChar, "\f", "\n")

		//regexp diye multiple space ba trim jodi thake easily remove kora jabe

		for _, singleLine := range strings.Split(replaceChar, "\n") {

			cleaningMultiSpace := re.ReplaceAllString(singleLine, " ")
			cleaningMultiSpace = strings.TrimSpace(cleaningMultiSpace)

			//multiple empty line check
			if cleaningMultiSpace == "" && lastLineEmpty {
				continue
			}
			if cleaningMultiSpace == "" {
				lastLineEmpty = true
			} else {
				lastLineEmpty = false
			}
			//empty space check
			trimmedLines = append(trimmedLines, cleaningMultiSpace)
		}
	}
	//SETEP 6 ENGIND
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

// Step 7: Transforming IATA & ICAO codes to Airport Names
func buildLookupMaps(data [][]string) (iataMap map[string]string, icaoMap map[string]string, cityMap map[string]string) {

	iataMap = make(map[string]string)
	icaoMap = make(map[string]string)
	cityMap = make(map[string]string)

	for _, row := range data[1:] { //[1:] header skip korar jonno
		name := row[0]
		icao := row[3]
		iata := row[4]
		city := row[2]

		iataMap[iata] = name //iata code diye khujle airportname paoa jabe
		icaoMap[icao] = name //	icao code diye khujle airportname paoa jabe
		cityMap[iata] = city //city name diye khujle airportname paoa jabe
		//ex bhalo er 3 ta english ache good,better,best. jeta search dibe ans ashbe bhalo.
	}
	return iataMap, icaoMap, cityMap

}

var pattern = `(#|\##|\*#)([A-Z]{3,4})` // golang er google docs a

func transformCodeToName(inputText []string, iataMap map[string]string, icaoMap map[string]string, cityMap map[string]string) []string {

	re := regexp.MustCompile(pattern)
	var outputLines []string
	for _, line := range inputText {
		lines := re.FindAllStringSubmatch(line, -1) //-1 means jotogulo match hoy shob dao. eita slice of slice return korbe. single string na. tai r ekta loop chalaite hobe

		for _, match := range lines {
			var result string
			fullMatch := match[0] //example: ##LAX
			prefix := match[1]    // ##
			code := match[2]      // LAX

			if prefix == "#" {
				result = iataMap[code]
			} else if prefix == "##" {
				result = icaoMap[code]
			} else if prefix == "*#" {
				result = cityMap[code]
			}
			if result != "" {
				line = strings.ReplaceAll(line, fullMatch, result) // line theke jeta match korse sheta replace kore dibe result diye
			}
		}
		outputLines = append(outputLines, line)
	}
	return outputLines

}
