package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
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

	if !validCSV(reader){
		os.Exit(1)
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
header := reader[0]
	for _, row := range reader {
		
		if len(header) != 6 {
			fmt.Fprintln(os.Stderr, "Airport lookup malformed")
			return false
		} 	
		if row == "" {
			fmt.Fprintln(os.Stderr, "Airport lookup malformed")
			return false
		}
	}

	for i = 0: i < 5; i++ {
		if header[i] != expectedHeader[i] {
			fmt.Fprintln(os.Stderr, "Airport lookup malformed")
			return false
	}	
	}
		
		return true
	}
	

