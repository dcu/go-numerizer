package numerizer

import (
	"strconv"
)

type NameValuePair struct {
	Name  string
	Value string
}

type NameValuePairs []*NameValuePair

func (pair *NameValuePair) ValueAsInt() int {
	value, err := strconv.Atoi(pair.Value)
	if err != nil {
		return 0.0
	}

	return value
}

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

	BIG_PREFIXES = NameValuePairs{
		&NameValuePair{"hundred", "100"},
		&NameValuePair{"thousand", "1000"},
		&NameValuePair{"million", "1000000"},
		&NameValuePair{"billion", "1000000000"},
		&NameValuePair{"trillion", "1000000000000"},
	}

	FRACTIONS = NameValuePairs{
		&NameValuePair{"half", "2"},
		&NameValuePair{"third(s)?", "3"},
		&NameValuePair{"fourth(s)?", "4"},
		&NameValuePair{"quarter(s)?", "4"},
		&NameValuePair{"fifth(s)?", "5"},
		&NameValuePair{"sixth(s)?", "6"},
		&NameValuePair{"seventh(s)?", "7"},
		&NameValuePair{"eighth(s)?", "8"},
		&NameValuePair{"nineth(s)?", "9"},
	}

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
