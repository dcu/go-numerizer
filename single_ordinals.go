package numerizer

import (
	"regexp"
)

var (
	SINGLE_ORDINALS = NameValuePairs{
		&NameValuePair{"first", "1"},
		&NameValuePair{"third", "3"},
		&NameValuePair{"fourth", "4"},
		&NameValuePair{"fifth", "5"},
		&NameValuePair{"sixth", "6"},
		&NameValuePair{"seventh", "7"},
		&NameValuePair{"eighth", "8"},
		&NameValuePair{"ninth", "9"},
	}
)

func replaceSingleOrdinals(text string) string {
	for _, singleOrdinalPair := range SINGLE_ORDINALS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + singleOrdinalPair.Name + `($|\W)`)
		suffix := singleOrdinalPair.Name[len(singleOrdinalPair.Name)-2:]
		text = rx.ReplaceAllString(text, `${1}<num>`+singleOrdinalPair.Value+suffix+`${2}`)
	}

	return text
}
