package numerizer

import (
	"strings"
)

func Numerize(text string) string {
	text = normalizeInput(text)
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

	return text
}
