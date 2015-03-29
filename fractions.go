package numerizer

import (
	"regexp"
	"strconv"
)

var (
	FRACTIONS = fractions{
		newFraction("half", "2"),
		newFraction("third(s)?", "3"),
		newFraction("fourth(s)?", "4"),
		newFraction("quarter(s)?", "4"),
		newFraction("fifth(s)?", "5"),
		newFraction("sixth(s)?", "6"),
		newFraction("seventh(s)?", "7"),
		newFraction("eighth(s)?", "8"),
		newFraction("nineth(s)?", "9"),
	}
)

var (
	evalFractionRx = regexp.MustCompile(`(?i)(\d+)(?: | and |-)+(<num>|\s)*(\d+)\s*\/\s*(\d+)`)
)

type fraction struct {
	name  string
	value string

	fractionRx1          *regexp.Regexp
	fractionReplacement1 string
	fractionRx2          *regexp.Regexp
	fractionReplacement2 string
}
type fractions []*fraction

func newFraction(name string, value string) *fraction {
	fraction := &fraction{name: name, value: value}

	fraction.fractionRx1 = regexp.MustCompile(`(?i)a ` + name + `(|\W)`)
	fraction.fractionReplacement1 = `<num>1/` + value + `${2}`

	fraction.fractionRx2 = regexp.MustCompile(`(?i)\s` + name + `($|\W)`)
	fraction.fractionReplacement2 = `/` + value + `${2}`
	return fraction
}

func replaceFractions(text string) string {
	for _, fractionPair := range FRACTIONS {
		text = fractionPair.fractionRx1.ReplaceAllString(text, fractionPair.fractionReplacement1)
		text = fractionPair.fractionRx2.ReplaceAllString(text, fractionPair.fractionReplacement2)
	}
	return text
}

func evaluateFractions(text string) string {
	// evaluate fractions when preceded by another number
	matches := evalFractionRx.FindStringSubmatch(text)
	if len(matches) == 5 {
		v1 := stringToFloat(matches[1])
		v3 := stringToFloat(matches[3])
		v4 := stringToFloat(matches[4])

		sum := v1 + (v3 / v4)

		replacement := strconv.FormatFloat(sum, 'f', -1, 64)

		text = evalFractionRx.ReplaceAllString(text, replacement)
	}

	return text
}
