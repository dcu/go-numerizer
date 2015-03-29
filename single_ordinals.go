package numerizer

import (
	"regexp"
)

var (
	SINGLE_ORDINALS = singleOrdinals{
		newSingleOrdinal("first", "1"),
		newSingleOrdinal("third", "3"),
		newSingleOrdinal("fourth", "4"),
		newSingleOrdinal("fifth", "5"),
		newSingleOrdinal("sixth", "6"),
		newSingleOrdinal("seventh", "7"),
		newSingleOrdinal("eighth", "8"),
		newSingleOrdinal("ninth", "9"),
	}
)

type singleOrdinal struct {
	name        string
	value       string
	suffix      string
	rx          *regexp.Regexp
	replacement string
}
type singleOrdinals []*singleOrdinal

func newSingleOrdinal(name string, value string) *singleOrdinal {
	pair := &singleOrdinal{name: name, value: value}
	pair.rx = regexp.MustCompile(`(?i)(^|\W)` + name + `($|\W)`)
	pair.suffix = name[len(name)-2:]
	pair.replacement = `${1}<num>` + value + pair.suffix + `${2}`

	return pair
}

func (pair *singleOrdinal) apply(text string) string {
	return pair.rx.ReplaceAllString(text, pair.replacement)
}

func replaceSingleOrdinals(text string) string {
	for _, singleOrdinalPair := range SINGLE_ORDINALS {
		text = singleOrdinalPair.apply(text)
	}

	return text
}
