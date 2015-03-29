package numerizer

import (
	"regexp"
	"strconv"
)

var (
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
)

func replaceFractions(text string) string {
	for _, fractionPair := range FRACTIONS {
		rx := regexp.MustCompile(`(?i)a ` + fractionPair.Name + `(|\W)`)
		text = rx.ReplaceAllString(text, `<num>1/`+fractionPair.Value+`${2}`)

		rx = regexp.MustCompile(`(?i)\s` + fractionPair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `/`+fractionPair.Value+`${2}`)
	}
	return text
}

func evaluateFractions(text string) string {
	// evaluate fractions when preceded by another number
	rx := regexp.MustCompile(`(?i)(\d+)(?: | and |-)+(<num>|\s)*(\d+)\s*\/\s*(\d+)`)
	matches := rx.FindStringSubmatch(text)
	if len(matches) == 5 {
		v1, _ := strconv.ParseFloat(matches[1], 64)
		v3, _ := strconv.ParseFloat(matches[3], 64)
		v4, _ := strconv.ParseFloat(matches[4], 64)

		sum := v1 + (v3 / v4)

		replacement := strconv.FormatFloat(sum, 'f', 2, 64)

		text = rx.ReplaceAllString(text, replacement)
	}

	return text
}
