package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
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

	iata_return, icao_return, city_return := buildLookupMaps(fileOutput)

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
			trimmedLines = append(trimmedLines, cleaningMultiSpace) // slice er jonno append use korte hoy. String er jonno noy.
		}
	}
	//SETEP 6 ENGIND

	//Step 7: Transforming IATA & ICAO codes to Airport Names
	transferResult := transformCodeToName(trimmedLines, iata_return, icao_return, city_return)

	//Step 8 : date Transform er jonno amar ekta alada function lagbe jeta transferResult theke paoa slice k range korle amra string pabw and oi string er moddhe date time pattern khujbo and dateTransform function call kore replace kore dibo.
	var finalResult []string
	for _, line := range transferResult {
		finalResult = append(finalResult, dateTransform(line)) //date transform er result final result e rakha
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

var pattern = `(\*#|##|#)([A-Z]{3,4})` // golang er google docs a

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
				result = cityMap[code] //fullmatch er first 2 char *# remove kore baki ta city code
			}
			if result != "" {
				line = strings.Replace(line, fullMatch, result, 1) // line theke jeta match korse sheta replace kore dibe result diye. 1 means first match ta replace korbe
			}
		}
		outputLines = append(outputLines, line)
	}
	return outputLines

}

var (
	timePattern       = regexp.MustCompile(`(D|T12|T24)\(([^)]+)\)`) //D(DD MMM YYYY) T12(HH:MM AM/PM) T24(HH:MM)
	offsetTimePattern = regexp.MustCompile(`(Z|[+-]\d{2}:\d{2})$`)
)

func dateTransform(inputTime string) string {
	catchTime := timePattern.FindAllStringSubmatch(inputTime, -1)
	workingStringTime := inputTime
	//RFC3339 format: 2006-01-02T15:04:05Z07:00 eita diye just emon format e support korbe but amder input onujai kokhono second nao thakte pare tai amader golan k format bujaiya dite hobe j ki ki format expected from input.

	//time er layout ekta slice a rakha
	timeLayout := []string{
		time.RFC3339, //2006-01-02T15:04:05Z07:00 ek e kotha// etar moddhe shob ache time second soho shob. but jodi second na thake tahole error dibe tai amader alada layout o dite hobe jate second na thakleo parse korte pare.
		// professional format holo boro format ta first a rakha
		"2006-01-02T15:04Z07:00", // without seconds

	}

	for _, match := range catchTime {

		fullDateMatch := match[0] //D(2007-04-05T12:30-02:00)
		tag := match[1]           //D|T12|T24
		rawValue := match[2]      //2007-04-05T12:30-02:00 (D case)	14:45 (T12 / T24 case)

		if !offsetTimePattern.MatchString(rawValue) {
			continue // offset pattern match na korle continue kore next match e jabe. karon amader expected input e date time er sheshe always offset thakbe. jodi na thake ba bhul thake tahole sheta valid input hobe na.
		}

		var parsedTime time.Time
		var err error
		success := false

		//for parsing
		for _, layout := range timeLayout {
			parsedTime, err = time.Parse(layout, rawValue) //raw value layout onujai parse korbe.
			if err == nil {
				success = true
				break // parsing successful hole loop theke ber hoye jabe
			}
		}
		//ei jaygata ektu tricky. success starting a var false and jodi err nil hoy tahole eita true and break kore loop theke ber hoye jabe r jodi !nil hoy tahole !success true hoye jabe and abar continue kore next match a khujte jabe. Its called Gaurd Clause.
		if !success {
			continue // jodi kono layout e parse na hoy tahole continue kore next match e jabe
		}
		//golang a time er formating er jonno rules : 01: মাস (Month) 02: দিন (Day)	03: ঘণ্টা (12h format) 04: মিনিট (Minute) 05: সেকেন্ড (Second) 06: বছর (Year) 07: টাইমজোন অফসেট (Timezone Offset)
		//offset finding and adjusting
		offsetString := parsedTime.Format("-07:00") // -07:00 fixed → Go numeric offset দেখায়, কোনো "Z" নয়

		var result string
		switch tag {
		case "D":
			result = parsedTime.Format("02 Jan 2006")

		case "T12":
			// parsedTime, err := time.Parse("15:04", rawValue) //time.parse (layout , value). input hisebe amader 15:00 ashte pare but output hisebe amader 03:00 PM dekhate hobe. tai layout e 15:04 dite hobe. karon time.parse er layout e 15:04 but output amar 12H format a dibe jokhon ami format set kore dibw tokhon.

			result = fmt.Sprintf("%s (%s)", parsedTime.Format("03:04PM"), offsetString)

		case "T24":

			result = fmt.Sprintf("%s (%s)", parsedTime.Format("15:04"), offsetString)
		}
		if result != "" {
			workingStringTime = strings.Replace(workingStringTime, fullDateMatch, result, 1) //input time theke jeta match korse sheta replace kore dibe result diye. 1 means first match ta replace korbe
		}
	}
	return workingStringTime

}
