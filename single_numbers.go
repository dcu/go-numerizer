package numerizer

import (
	"regexp"
)

var (
	SINGLE_NUMS = singleNumbers{
		newSingleNumber("one", "1"),
		newSingleNumber("two", "2"),
		newSingleNumber("three", "3"),
		newSingleNumber("four", "4"),
		newSingleNumber("five", "5"),
		newSingleNumber("six", "6"),
		newSingleNumber("seven", "7"),
		newSingleNumber("eight", "8"),
		newSingleNumber("nine", "9"),
	}
)

type singleNumber struct {
	name        string
	value       string
	rx          *regexp.Regexp
	replacement string
}
type singleNumbers []*singleNumber

func (pair *singleNumber) apply(text string) string {
	return pair.rx.ReplaceAllString(text, pair.replacement)
}

func newSingleNumber(name string, value string) *singleNumber {
	pair := &singleNumber{name: name, value: value}

	pair.rx = regexp.MustCompile(`(?i)(^|\W)` + name + `($|\W)`)
	pair.replacement = `${1}<num>` + value + `${2}`

	return pair
}

func replaceSingleNumbers(text string) string {
	for _, pair := range SINGLE_NUMS {
		text = pair.apply(text)
	}

	return text
}
