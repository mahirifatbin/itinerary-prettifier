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
func main() {

	flag.Parse()

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
		fmt.Fprint(os.Stderr, "Airport lookup malformed")
		os.Exit(1)
	}

	//Input file reading
	fileInput, err := filehandle.InputTextRead(input)
	if err != nil {
		fmt.Fprint(os.Stderr, "Input not found")
		os.Exit(1)
	}

	//STEP 6 : input text clean kora
	trimmedLines := clean.TrimWhitespace(fileInput) //trim or clean kora input data store kora

	// Step 7: Transforming IATA & ICAO codes to Airport Names
	iata_return, icao_return, city_return := csvlookup.BuildLookupMaps(fileOutput)

	//Step 7: Transforming IATA & ICAO codes to Airport Names
	transferResult := transform.TransformCodeToName(trimmedLines, iata_return, icao_return, city_return)

	//Step 8 : date Transform er jonno amar ekta alada function lagbe jeta transferResult theke paoa slice k range korle amra string pabw and oi string er moddhe date time pattern khujbo and dateTransform function call kore replace kore dibo.
	var finalResult []string
	for _, line := range transferResult {
		finalResult = append(finalResult, transform.DateTransform(line)) //date transform er result final result e rakha
	}

	//output file writing
	fileOut, err := os.Create(output)
	if err != nil {
		fmt.Fprint(os.Stderr, "Cannot create output file\n")
		os.Exit(1)
	}
	defer fileOut.Close()
	writingResult := bufio.NewWriter(fileOut)
	for _, outputResultWriting := range finalResult {

		writingResult.WriteString(outputResultWriting + "\n") //WriteString Ram a joma kore rakhe
	}
	writingResult.Flush() // RAM a joma kora data disk e store kore

	//important: main function a main kokhono kichu return korbe nah.

}
