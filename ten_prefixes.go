package numerizer

import (
	"regexp"
	"strconv"
)

var (
	TEN_PREFIXES = NameValuePairs{
		&NameValuePair{"twenty", "20"},
		&NameValuePair{"thirty", "30"},
		&NameValuePair{"forty", "40"},
		&NameValuePair{"fourty", "40"}, // Common misspelling
		&NameValuePair{"fifty", "50"},
		&NameValuePair{"sixty", "60"},
		&NameValuePair{"seventy", "70"},
		&NameValuePair{"eighty", "80"},
		&NameValuePair{"ninety", "90"},
	}
)

func replaceTenPrefixes(text string) string {
	for _, tenPrefixPair := range TEN_PREFIXES {
		for _, singleNumPair := range SINGLE_NUMS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixPair.Name + singleNumPair.Name + `($|\W)`)

			num := strconv.Itoa(tenPrefixPair.ValueAsInt() + singleNumPair.ValueAsInt())

			text = rx.ReplaceAllString(text, `${1}<num>`+num+`${2}`)
		}

		for _, singleOrdinalPair := range SINGLE_ORDINALS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixPair.Name + `(\s)?` + singleOrdinalPair.Name + `($|\W)`)

			num := strconv.Itoa(tenPrefixPair.ValueAsInt() + singleOrdinalPair.ValueAsInt())
			suffix := singleOrdinalPair.Name[len(singleOrdinalPair.Name)-2:]

			text = rx.ReplaceAllString(text, `${1}<num>`+num+suffix+`${3}`)
		}

		rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixPair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+tenPrefixPair.Value+`${2}`)
	}
	return text
}
