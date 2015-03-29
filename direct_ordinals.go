package numerizer

import (
	"regexp"
)

var (
	DIRECT_ORDINALS = directOrdinals{
		newDirectOrdinal("tenth", "10"),
		newDirectOrdinal("eleventh", "11"),
		newDirectOrdinal("twelfth", "12"),
		newDirectOrdinal("thirteenth", "13"),
		newDirectOrdinal("fourteenth", "14"),
		newDirectOrdinal("fifteenth", "15"),
		newDirectOrdinal("sixteenth", "16"),
		newDirectOrdinal("seventeenth", "17"),
		newDirectOrdinal("eighteenth", "18"),
		newDirectOrdinal("nineteenth", "19"),
		newDirectOrdinal("twentieth", "20"),
		newDirectOrdinal("thirtieth", "30"),
		newDirectOrdinal("fourtieth", "40"),
		newDirectOrdinal("fiftieth", "50"),
		newDirectOrdinal("sixtieth", "60"),
		newDirectOrdinal("seventieth", "70"),
		newDirectOrdinal("eightieth", "80"),
		newDirectOrdinal("ninetieth", "90"),
	}
)

type directOrdinal struct {
	name        string
	value       string
	suffix      string
	rx          *regexp.Regexp
	replacement string
}
type directOrdinals []directOrdinal

func newDirectOrdinal(name string, value string) directOrdinal {
	pair := directOrdinal{name: name, value: value}
	pair.rx = regexp.MustCompile(`(?i)(^|\W)` + name + `($|\W)`)
	pair.suffix = name[len(name)-2:]
	pair.replacement = `${1}<num>` + pair.value + pair.suffix + `${2}`

	return pair
}

func (pair directOrdinal) apply(text string) string {
	text = pair.rx.ReplaceAllString(text, pair.replacement)

	return text
}

func replaceDirectOrdinals(text string) string {
	for _, directOrdinalPair := range DIRECT_ORDINALS {
		text = directOrdinalPair.apply(text)
	}

	return text
}
