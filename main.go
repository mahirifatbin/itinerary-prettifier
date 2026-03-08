package main

import (
	"Itinerary/internal/clean"
	"Itinerary/internal/csvlookup"
	"Itinerary/internal/filehandle"
	"Itinerary/internal/transform"
	"bufio"
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
func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `itinerary usage:
go run . ./input.txt ./output.txt ./airport-lookup.csv`)
	}
}

func main() {

	flagVersion := flag.Bool("v", false, "Print version information and exit")                           // version flag
	compressFlag := flag.Bool("c", false, "Compress multiple spaces and tabs into a single space space") // compress flag for horizontal space compression

	flag.Parse()
	if *flagVersion {
		fmt.Println("itinerary version 1.0.0")
		os.Exit(0)
	}

	if flag.NArg() != 3 {
		flag.Usage()
		os.Exit(1)
	}

	input := flag.Arg(0)
	output := flag.Arg(1)
	airportLookup := flag.Arg(2)

	filehandle.CheckFileExist(input, airportLookup)

	//CSV file opening and validation
	fileOutput, err := filehandle.OpenCSVFile(airportLookup)
	if err != nil {
		fmt.Fprint(os.Stderr, "Airport lookup not found")
		os.Exit(1)
	}

	if !csvlookup.ValidCSV(fileOutput) {
		fmt.Fprint(os.Stderr, "Airport lookup malformed\n")
		os.Exit(1)
	}

	//Input file reading
	fileInput, err := filehandle.InputTextRead(input)
	if err != nil {
		fmt.Fprint(os.Stderr, "Input not found")
		os.Exit(1)
	}

	//Input text cleaning and trimming
	trimmedLines := clean.TrimWhitespace(fileInput, *compressFlag) //trim or clean kora input data store kora

	//Transforming IATA & ICAO codes to Airport Names
	iata_return, icao_return, city_return := csvlookup.BuildLookupMaps(fileOutput)

	//Transforming IATA & ICAO codes to Airport Names
	transferResult := transform.TransformCodeToName(trimmedLines, iata_return, icao_return, city_return)

	//Transforming Date and Time Formats
	var finalResult []string
	for _, line := range transferResult {
		finalResult = append(finalResult, transform.DateTransform(line)) //date transform results in final result.
	}

	//Output file writing
	fileOut, err := os.Create(output)
	if err != nil {
		fmt.Fprint(os.Stderr, "Cannot create output file\n")
		os.Exit(1)
	}
	defer fileOut.Close()

	//default Flight Information output heading
	header := "✈ Flight Information\n--------------------\n\n"
	fileOut.WriteString(header)

	writingResult := bufio.NewWriter(fileOut)
	for _, outputResultWriting := range finalResult {

		writingResult.WriteString(outputResultWriting + "\n") //WriteString store in RAM
	}
	writingResult.Flush() // RAM data stored in Disk.

}
