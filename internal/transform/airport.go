package transform

import (
	"regexp"
	"strings"
)

var pattern = `(\*#|##|#)([A-Z]{3,4})`

func TransformCodeToName(inputText []string, iataMap map[string]string, icaoMap map[string]string, cityMap map[string]string) []string {

	re := regexp.MustCompile(pattern)
	var outputLines []string
	for _, line := range inputText {
		lines := re.FindAllStringSubmatch(line, -1) //-1 means getting all match.

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
				line = strings.Replace(line, fullMatch, result, 1) // 1 means first match will be replaced
			}
		}
		outputLines = append(outputLines, line)
	}
	return outputLines

}
