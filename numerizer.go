package numerizer

import (
	"regexp"
	"strings"
)

var (
	HyphenatedWordsRx = regexp.MustCompile(` +|([^\d])-([^\d])`)
	AnditionRx        = regexp.MustCompile(`(?i)<num>(\d+)( | and )<num>(\d+)([^\w]|$)`)
)

func Numerize(text string) (string, error) {
	text = HyphenatedWordsRx.ReplaceAllString(text, "${1} ${2}")
	var err error

	text = replaceDirectNumbers(text)
	text = replaceSingleNumbers(text)
	text = replaceTenPrefixes(text)
	text = replaceFractions(text)
	text = replaceDirectOrdinals(text)
	text = replaceSingleOrdinals(text)
	text = evaluateFractions(text)
	text = replaceBigPrefixes(text)

	text = andition(text)
	text = strings.Replace(text, `<num>`, "", -1)

	return text, err
}
