package numerizer

import (
	"regexp"
)

var (
	SINGLE_NUMS = NameValuePairs{
		&NameValuePair{"one", "1"},
		&NameValuePair{"two", "2"},
		&NameValuePair{"three", "3"},
		&NameValuePair{"four", "4"},
		&NameValuePair{"five", "5"},
		&NameValuePair{"six", "6"},
		&NameValuePair{"seven", "7"},
		&NameValuePair{"eight", "8"},
		&NameValuePair{"nine", "9"},
	}
)

func replaceSingleNumbers(text string) string {
	for _, pair := range SINGLE_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + pair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+pair.Value+`${2}`)
	}

	return text
}
