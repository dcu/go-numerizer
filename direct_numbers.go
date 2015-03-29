package numerizer

import (
	"regexp"
)

var (
	DIRECT_NUMS = directNumbers{
		newDirectNumber("eleven", "11"),
		newDirectNumber("twelve", "12"),
		newDirectNumber("thirteen", "13"),
		newDirectNumber("fourteen", "14"),
		newDirectNumber("fifteen", "15"),
		newDirectNumber("sixteen", "16"),
		newDirectNumber("seventeen", "17"),
		newDirectNumber("eighteen", "18"),
		newDirectNumber("nineteen", "19"),
		newDirectNumber("ninteen", "19"), // Common mis-spelling
		newDirectNumber("zero", "0"),
		newDirectNumber("ten", "10"),
		newDirectNumber("\ba[\b^$]", "1"), // doesn't make sense for an 'a' at the end to be a 1
	}
)

type directNumber struct {
	Name        string
	Value       string
	rx          *regexp.Regexp
	replacement string
}

type directNumbers []*directNumber

func newDirectNumber(name string, value string) *directNumber {
	pair := &directNumber{Name: name, Value: value}

	pair.rx = regexp.MustCompile(`(?i)(^|\W)` + name + `($|\W)`)
	pair.replacement = `${1}<num>` + pair.Value + `${2}`

	return pair
}

func (pair *directNumber) apply(text string) string {
	return pair.rx.ReplaceAllString(text, pair.replacement)
}

func replaceDirectNumbers(text string) string {
	for _, pair := range DIRECT_NUMS {
		text = pair.apply(text)
	}

	return text
}
