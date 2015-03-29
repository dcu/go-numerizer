package numerizer

import (
	"regexp"
	"strconv"
)

var (
	anditionRx        = regexp.MustCompile(`(?i)<num>(\d+)( | and )<num>(\d+)([^\w]|$)`)
	hyphenatedWordsRx = regexp.MustCompile(` +|([^\d])-([^\d])`)
)

// Converts "20 7" or "20 and 7" to "27"
func andition(text string) string {
	for {
		allMatches := anditionRx.FindStringSubmatch(text)
		if len(allMatches) != 5 {
			break
		}

		v1 := stringToInt(allMatches[1])
		v2 := stringToInt(allMatches[3])

		if allMatches[1] == "and" || len(allMatches[1]) > len(allMatches[3]) {
			total := v1 + v2
			totalStr := strconv.Itoa(total)
			replacement := `<num>` + totalStr + allMatches[4]

			text = replaceOnce(anditionRx, text, replacement)
		}
	}

	return text
}

func replaceOnce(rx *regexp.Regexp, text string, replacement string) string {
	wasReplaced := false
	text = rx.ReplaceAllStringFunc(text, func(src string) string {
		if !wasReplaced {
			src = replacement
		}

		wasReplaced = true
		return src
	})

	return text
}

func stringToInt(integerStr string) int {
	value, err := strconv.Atoi(integerStr)
	if err != nil {
		value = 0
	}

	return value
}

func stringToFloat(floatStr string) float64 {
	value, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		value = 0.0
	}

	return value
}

func normalizeInput(input string) string {
	input = hyphenatedWordsRx.ReplaceAllString(input, "${1} ${2}")

	return input
}
