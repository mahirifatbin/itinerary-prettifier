package transform

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	timePattern       = regexp.MustCompile(`(D|T12|T24)\(([^)]+)\)`) //D(DD MMM YYYY) T12(HH:MM AM/PM) T24(HH:MM)
	offsetTimePattern = regexp.MustCompile(`(Z|[+-]\d{2}:\d{2})$`)
)

func DateTransform(inputTime string) string {
	catchTime := timePattern.FindAllStringSubmatch(inputTime, -1)
	workingStringTime := inputTime

	//Slice in a time layout.
	timeLayout := []string{
		time.RFC3339,             //2006-01-02T15:04:05Z07:00 (ISO 8601 format with seconds)
		"2006-01-02T15:04Z07:00", // without seconds

	}

	for _, match := range catchTime {

		fullDateMatch := match[0] //D(2007-04-05T12:30-02:00)
		tag := match[1]           //D|T12|T24
		rawValue := match[2]      //2007-04-05T12:30-02:00 (D case)	14:45 (T12 / T24 case)

		if !offsetTimePattern.MatchString(rawValue) {
			fmt.Fprintf(os.Stderr, "[WARN] Skipping invalid offset format: %s\n", rawValue)
			continue
		}

		var parsedTime time.Time
		var err error
		success := false

		//for parsing
		for _, layout := range timeLayout {
			parsedTime, err = time.Parse(layout, rawValue) //raw value will parsed based on layout.
			if err == nil {
				success = true
				break // parsing successful then get out of the loop
			}
		}

		if !success {
			fmt.Fprintf(os.Stderr, "[WARN] Skipping invalid date/time input: %s\n", fullDateMatch)
			continue // if no parse in any layout then skip this match and continue with next match.

		}

		//offset finding and adjusting
		offsetString := parsedTime.Format("-07:00")

		var result string
		switch tag {
		case "D":
			result = parsedTime.Format("02 Jan 2006")

		case "T12":

			result = fmt.Sprintf("%s (%s)", parsedTime.Format("03:04PM"), offsetString)

		case "T24":

			result = fmt.Sprintf("%s (%s)", parsedTime.Format("15:04"), offsetString)
		}
		if result != "" {
			workingStringTime = strings.Replace(workingStringTime, fullDateMatch, result, 1) //1 means first match will be replaced
		}
	}
	return workingStringTime

}
