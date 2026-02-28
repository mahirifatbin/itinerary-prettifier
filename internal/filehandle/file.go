package filehandle

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func CheckFileExist(input string, airportLookup string) {
	_, err := os.Stat(input)
	_, err1 := os.Stat(airportLookup)

	if os.IsNotExist(err) {
		fmt.Fprint(os.Stderr, "Input not found\n")
		os.Exit(1)
	}
	if os.IsNotExist(err1) {
		fmt.Fprint(os.Stderr, "Airport lookup not found\n")
		os.Exit(1)
	}
}

func OpenCSVFile(path string) ([][]string, error) {
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

//input text reading

func InputTextRead(path string) ([]string, error) {

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
