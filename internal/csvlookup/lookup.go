package csvlookup

import (
	"strings"
)

var requiredColumnHeader = []string{
	"name",
	"iso_country",
	"municipality",
	"icao_code",
	"iata_code",
	"coordinates",
}

var columnIndex map[string]int //header map globally store করতে হবে

// csv validation
func ValidCSV(reader [][]string) bool {
	//header validation
	if len(reader) == 0 {
		return false
	}
	header := reader[0]
	if len(header) != 6 {

		return false
	}

	//header map create
	columnIndex = make(map[string]int)

	for i, column := range header {
		columnIndex[column] = i
	}
	for _, col := range requiredColumnHeader {
		if _, ok := columnIndex[col]; !ok {
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

// Step 7: Transforming IATA & ICAO codes to Airport Names
func BuildLookupMaps(data [][]string) (iataMap map[string]string, icaoMap map[string]string, cityMap map[string]string) {

	iataMap = make(map[string]string)
	icaoMap = make(map[string]string)
	cityMap = make(map[string]string)

	for _, row := range data[1:] { //[1:] header skip korar jonno
		name := row[columnIndex["name"]]
		icao := row[columnIndex["icao_code"]]
		iata := row[columnIndex["iata_code"]]
		city := row[columnIndex["municipality"]]

		iataMap[iata] = name //iata code diye khujle airportname paoa jabe
		icaoMap[icao] = name //	icao code diye khujle airportname paoa jabe
		cityMap[iata] = city //city name diye khujle airportname paoa jabe
		//ex bhalo er 3 ta english ache good,better,best. jeta search dibe ans ashbe bhalo.
	}
	return iataMap, icaoMap, cityMap

}
