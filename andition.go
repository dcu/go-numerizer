package numerizer

import (
	"strconv"
)

// Converts "20 7" or "20 and 7" to "27"
func andition(text string) string {
	count := 0
	for {
		allMatches := AnditionRx.FindStringSubmatch(text)

		if len(allMatches) != 5 {
			break
		}

		v1, err := strconv.Atoi(allMatches[1])
		if err != nil {
			break
		}
		v2, err := strconv.Atoi(allMatches[3])
		if err != nil {
			break
		}

		if allMatches[1] == "and" || len(allMatches[1]) > len(allMatches[3]) {
			total := v1 + v2
			totalStr := strconv.Itoa(total)
			replacement := `<num>` + totalStr + allMatches[4]
			replacementCount := 0
			text = AnditionRx.ReplaceAllStringFunc(text, func(src string) string {
				if replacementCount == 0 {
					src = replacement
				}

				replacementCount += 1

				return src
			})
		}

		count += 1
	}

	return text
}
