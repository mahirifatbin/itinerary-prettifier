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
			fmt.Fprintf(os.Stderr, "[WARN] Skipping invalid offset format: %s\n", rawValue)
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
			fmt.Fprintf(os.Stderr, "[WARN] Skipping invalid date/time input: %s\n", fullDateMatch)
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
