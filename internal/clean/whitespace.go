package clean

import (
	"regexp"
	"strings"
)

func TrimWhitespace(fileInput []string) []string {
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
	return trimmedLines
}
