package numerizer

import (
	"regexp"
	"strconv"
)

var (
	BIG_PREFIXES = bigPrefixes{
		NewBigPrefix("hundred", "100"),
		NewBigPrefix("thousand", "1000"),
		NewBigPrefix("million", "1000000"),
		NewBigPrefix("billion", "1000000000"),
		NewBigPrefix("trillion", "1000000000000"),
	}
)

type bigPrefix struct {
	name       string
	value      string
	valueAsInt int
	rx         *regexp.Regexp
}
type bigPrefixes []*bigPrefix

func NewBigPrefix(name string, value string) *bigPrefix {
	bigPrefixPair := &bigPrefix{name: name, value: value}
	bigPrefixPair.rx = regexp.MustCompile(`(?i)(?:<num>)?(\d*) *` + name)
	bigPrefixPair.valueAsInt = stringToInt(value)

	return bigPrefixPair
}

func (pair *bigPrefix) apply(text string) string {
	matches := pair.rx.FindStringSubmatch(text)

	if len(matches) == 2 {
		replacement := pair.value
		if len(matches[1]) > 0 {
			v1 := stringToInt(matches[1])
			replacement = `<num>` + strconv.Itoa(v1*pair.valueAsInt)
		}

		text = pair.rx.ReplaceAllString(text, replacement)
	}

	return andition(text)
}

// hundreds, thousands, millions, etc.
func replaceBigPrefixes(text string) string {
	for _, pair := range BIG_PREFIXES {
		text = pair.apply(text)
	}

	return text
}
