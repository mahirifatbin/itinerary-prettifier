package clean

import (
	"regexp"
	"strings"
)

func TrimWhitespace(fileInput []string, compressHorizontalSpace bool) []string {
	var trimmedLines []string // this will store new cleaned lines
	lastLineEmpty := false

	var re *regexp.Regexp
	if compressHorizontalSpace {
		re = regexp.MustCompile(`[ \t]+`) //multiple space or tab will remove and replace with single space
	}

	//FileInput is the slice of string where each element is a line of the input file. So we will iterate through each line and clean it.
	for _, line := range fileInput {

		trimmed := strings.TrimSpace(line) //Forward and backward space remove

		replaceChar := trimmed
		replaceChar = strings.ReplaceAll(replaceChar, "\r", "\n")
		replaceChar = strings.ReplaceAll(replaceChar, "\v", "\n")
		replaceChar = strings.ReplaceAll(replaceChar, "\f", "\n")

		for _, singleLine := range strings.Split(replaceChar, "\n") {

			cleaningMultiSpace := singleLine
			if compressHorizontalSpace && re != nil {
				cleaningMultiSpace = re.ReplaceAllString(cleaningMultiSpace, " ")
			}
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
			trimmedLines = append(trimmedLines, cleaningMultiSpace) // Append is using for adding new element in slice. Here we are adding cleaned line in trimmedLines slice.
		}

	}
	return trimmedLines
}
