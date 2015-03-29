package numerizer

import (
	"regexp"
)

var (
	DIRECT_ORDINALS = NameValuePairs{
		&NameValuePair{"tenth", "10"},
		&NameValuePair{"eleventh", "11"},
		&NameValuePair{"twelfth", "12"},
		&NameValuePair{"thirteenth", "13"},
		&NameValuePair{"fourteenth", "14"},
		&NameValuePair{"fifteenth", "15"},
		&NameValuePair{"sixteenth", "16"},
		&NameValuePair{"seventeenth", "17"},
		&NameValuePair{"eighteenth", "18"},
		&NameValuePair{"nineteenth", "19"},
		&NameValuePair{"twentieth", "20"},
		&NameValuePair{"thirtieth", "30"},
		&NameValuePair{"fourtieth", "40"},
		&NameValuePair{"fiftieth", "50"},
		&NameValuePair{"sixtieth", "60"},
		&NameValuePair{"seventieth", "70"},
		&NameValuePair{"eightieth", "80"},
		&NameValuePair{"ninetieth", "90"},
	}
)

func replaceDirectOrdinals(text string) string {
	for _, directOrdinalPair := range DIRECT_ORDINALS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + directOrdinalPair.Name + `($|\W)`)
		suffix := directOrdinalPair.Name[len(directOrdinalPair.Name)-2:]
		text = rx.ReplaceAllString(text, `${1}<num>`+directOrdinalPair.Value+suffix+`${2}`)
	}

	return text
}
