package numerizer

import (
	"regexp"
	"strconv"
)

var (
	BIG_PREFIXES = NameValuePairs{
		&NameValuePair{"hundred", "100"},
		&NameValuePair{"thousand", "1000"},
		&NameValuePair{"million", "1000000"},
		&NameValuePair{"billion", "1000000000"},
		&NameValuePair{"trillion", "1000000000000"},
	}
)

func replaceBigPrefixes(text string) string {
	// hundreds, thousands, millions, etc.
	for _, pair := range BIG_PREFIXES {
		rx := regexp.MustCompile(`(?i)(?:<num>)?(\d*) *` + pair.Name)
		matches := rx.FindStringSubmatch(text)

		if len(matches) == 2 {
			replacement := pair.Value
			if len(matches[1]) > 0 {
				v1 := stringToInt(matches[1])
				replacement = `<num>` + strconv.Itoa(v1*pair.ValueAsInt())
			}

			text = rx.ReplaceAllString(text, replacement)
		}

		text = andition(text)
	}

	return text
}
