package numerizer

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	HyphenatedWordsRx = regexp.MustCompile(` +|([^\d])-([^\d])`)
	AnditionRx        = regexp.MustCompile(`(?i)<num>(\d+)( | and )<num>(\d+)([^\w]|$)`)
)

func Numerize(text string) (string, error) {
	text = HyphenatedWordsRx.ReplaceAllString(text, "${1} ${2}")
	var err error

	for _, pair := range DIRECT_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + pair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+pair.Value+`${2}`)
	}

	for _, pair := range SINGLE_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + pair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+pair.Value+`${2}`)
	}

	for _, tenPrefixPair := range TEN_PREFIXES {
		for _, singleNumPair := range SINGLE_NUMS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixPair.Name + singleNumPair.Name + `($|\W)`)

			num := strconv.Itoa(tenPrefixPair.ValueAsInt() + singleNumPair.ValueAsInt())

			text = rx.ReplaceAllString(text, `${1}<num>`+num+`${2}`)
		}

		for _, singleOrdinalPair := range SINGLE_ORDINALS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixPair.Name + `(\s)?` + singleOrdinalPair.Name + `($|\W)`)

			num := strconv.Itoa(tenPrefixPair.ValueAsInt() + singleOrdinalPair.ValueAsInt())
			suffix := singleOrdinalPair.Name[len(singleOrdinalPair.Name)-2:]

			text = rx.ReplaceAllString(text, `${1}<num>`+num+suffix+`${3}`)
		}

		rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixPair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+tenPrefixPair.Value+`${2}`)
	}

	for _, fractionPair := range FRACTIONS {
		rx := regexp.MustCompile(`(?i)a ` + fractionPair.Name + `(|\W)`)
		text = rx.ReplaceAllString(text, `<num>1/`+fractionPair.Value+`${2}`)

		rx = regexp.MustCompile(`(?i)\s` + fractionPair.Name + `($|\W)`)
		text = rx.ReplaceAllString(text, `/`+fractionPair.Value+`${2}`)
	}

	for _, directOrdinalPair := range DIRECT_ORDINALS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + directOrdinalPair.Name + `($|\W)`)
		suffix := directOrdinalPair.Name[len(directOrdinalPair.Name)-2:]
		text = rx.ReplaceAllString(text, `${1}<num>`+directOrdinalPair.Value+suffix+`${2}`)
	}

	for _, singleOrdinalPair := range SINGLE_ORDINALS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + singleOrdinalPair.Name + `($|\W)`)
		suffix := singleOrdinalPair.Name[len(singleOrdinalPair.Name)-2:]
		text = rx.ReplaceAllString(text, `${1}<num>`+singleOrdinalPair.Value+suffix+`${2}`)
	}

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

	// hundreds, thousands, millions, etc.
	for _, pair := range BIG_PREFIXES {
		rx := regexp.MustCompile(`(?i)(?:<num>)?(\d*) *` + pair.Name)
		matches := rx.FindStringSubmatch(text)

		if len(matches) == 2 {
			replacement := pair.Value
			if len(matches[1]) > 0 {
				v1, _ := strconv.Atoi(matches[1])
				replacement = `<num>` + strconv.Itoa(v1*pair.ValueAsInt())
			}

			text = rx.ReplaceAllString(text, replacement)
		}

		text = andition(text)
	}

	text = andition(text)
	text = strings.Replace(text, `<num>`, "", -1)

	return text, err
}

// Converts "20 7" or "20 and 7" to "27"
func andition(text string) string {
	count := 0
	for {
		allMatches := AnditionRx.FindStringSubmatch(text)

		if len(allMatches) != 5 {
			break
		}

		v1, err := strconv.Atoi(allMatches[1])
		if err != nil {
			break
		}
		v2, err := strconv.Atoi(allMatches[3])
		if err != nil {
			break
		}

		if allMatches[1] == "and" || len(allMatches[1]) > len(allMatches[3]) {
			total := v1 + v2
			totalStr := strconv.Itoa(total)
			replacement := `<num>` + totalStr + allMatches[4]
			replacementCount := 0
			text = AnditionRx.ReplaceAllStringFunc(text, func(src string) string {
				if replacementCount == 0 {
					src = replacement
				}

				replacementCount += 1

				return src
			})
		}

		count += 1
	}

	return text
}
