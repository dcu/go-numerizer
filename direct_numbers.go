package numerizer

import (
	"regexp"
)

var (
	DIRECT_NUMS = NameValuePairs{
		&NameValuePair{"eleven", "11"},
		&NameValuePair{"twelve", "12"},
		&NameValuePair{"thirteen", "13"},
		&NameValuePair{"fourteen", "14"},
		&NameValuePair{"fifteen", "15"},
		&NameValuePair{"sixteen", "16"},
		&NameValuePair{"seventeen", "17"},
		&NameValuePair{"eighteen", "18"},
		&NameValuePair{"nineteen", "19"},
		&NameValuePair{"ninteen", "19"}, // Common mis-spelling
		&NameValuePair{"zero", "0"},
		&NameValuePair{"ten", "10"},
		&NameValuePair{"\ba[\b^$]", "1"}, // doesn't make sense for an 'a' at the end to be a 1
	}
)

func replaceDirectNumbers(text string) string {
	for _, pair := range DIRECT_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + pair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+pair.Value+`${2}`)
	}

	return text
}
