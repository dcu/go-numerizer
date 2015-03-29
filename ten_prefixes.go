package numerizer

import (
	"regexp"
	"strconv"
)

var (
	TEN_PREFIXES = tenPrefixes{
		newTenPrefix("twenty", "20"),
		newTenPrefix("thirty", "30"),
		newTenPrefix("forty", "40"),
		newTenPrefix("fourty", "40"), // Common misspelling
		newTenPrefix("fifty", "50"),
		newTenPrefix("sixty", "60"),
		newTenPrefix("seventy", "70"),
		newTenPrefix("eighty", "80"),
		newTenPrefix("ninety", "90"),
	}
)

type rxReplacement struct {
	rx          *regexp.Regexp
	replacement string
}
type rxReplacements []*rxReplacement

type tenPrefix struct {
	name        string
	value       string
	rx          *regexp.Regexp
	replacement string

	singleNumReplacements     rxReplacements
	singleOrdinalReplacements rxReplacements
}
type tenPrefixes []*tenPrefix

func newTenPrefix(name string, value string) *tenPrefix {
	pair := &tenPrefix{name: name, value: value}

	for _, singleNumPair := range SINGLE_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + name + singleNumPair.name + `($|\W)`)
		num := strconv.Itoa(stringToInt(value) + stringToInt(singleNumPair.value))
		replacement := `${1}<num>` + num + `${2}`

		pair.singleNumReplacements = append(pair.singleNumReplacements, &rxReplacement{rx, replacement})
	}

	for _, singleOrdinalPair := range SINGLE_ORDINALS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + name + `(\s)?` + singleOrdinalPair.name + `($|\W)`)
		num := strconv.Itoa(stringToInt(value) + stringToInt(singleOrdinalPair.value))
		replacement := `${1}<num>` + num + singleOrdinalPair.suffix + `${3}`

		pair.singleOrdinalReplacements = append(pair.singleOrdinalReplacements, &rxReplacement{rx, replacement})
	}

	pair.rx = regexp.MustCompile(`(?i)(^|\W)` + name + `($|\W)`)
	pair.replacement = `${1}<num>` + value + `${2}`

	return pair
}

func (pair *tenPrefix) apply(text string) string {
	for _, singleNumReplacement := range pair.singleNumReplacements {
		text = singleNumReplacement.rx.ReplaceAllString(text, singleNumReplacement.replacement)
	}

	for _, singleOrdinalReplacement := range pair.singleOrdinalReplacements {
		text = singleOrdinalReplacement.rx.ReplaceAllString(text, singleOrdinalReplacement.replacement)
	}
	text = pair.rx.ReplaceAllString(text, pair.replacement)

	return text
}

func replaceTenPrefixes(text string) string {
	for _, tenPrefixPair := range TEN_PREFIXES {
		text = tenPrefixPair.apply(text)
	}
	return text
}
