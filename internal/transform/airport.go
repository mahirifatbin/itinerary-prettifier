package transform

import (
	"regexp"
	"strings"
)

var pattern = `(\*#|##|#)([A-Z]{3,4})` // golang er google docs a

func TransformCodeToName(inputText []string, iataMap map[string]string, icaoMap map[string]string, cityMap map[string]string) []string {

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
