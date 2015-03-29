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

	for key, value := range DIRECT_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + key + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+value+`${2}`)
	}

	for key, value := range SINGLE_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + key + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+value+`${2}`)
	}

	for tenPrefixName, tenPrefixStr := range TEN_PREFIXES {
		for singleNumName, singleNumStr := range SINGLE_NUMS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixName + singleNumName + `($|\W)`)

			tenPrefixNum, _ := strconv.Atoi(tenPrefixStr)
			singleNumNum, _ := strconv.Atoi(singleNumStr)

			num := strconv.Itoa(tenPrefixNum + singleNumNum)

			text = rx.ReplaceAllString(text, `${1}<num>`+num+`${2}`)
		}

		for singleOrdinalName, singleOrdinalStr := range SINGLE_ORDINALS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixName + `(\s)?` + singleOrdinalName + `($|\W)`)

			tenPrefixNum, _ := strconv.Atoi(tenPrefixStr)
			singleOrdinalNum, _ := strconv.Atoi(singleOrdinalStr)

			num := strconv.Itoa(tenPrefixNum + singleOrdinalNum)
			suffix := singleOrdinalName[len(singleOrdinalName)-2:]

			text = rx.ReplaceAllString(text, `${1}<num>`+num+suffix+`${3}`)
		}

		rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixName + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+tenPrefixStr+`${2}`)
	}

	for fractionName, fractionValueStr := range FRACTIONS {
		rx := regexp.MustCompile(`(?i)a ` + fractionName + `(|\W)`)
		text = rx.ReplaceAllString(text, `<num>1/`+fractionValueStr+`${2}`)

		rx = regexp.MustCompile(`(?i)\s` + fractionName + `($|\W)`)
		text = rx.ReplaceAllString(text, `/`+fractionValueStr+`${2}`)
	}

	for directOrdinalName, directOrdinalStr := range DIRECT_ORDINALS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + directOrdinalName + `($|\W)`)
		suffix := directOrdinalName[len(directOrdinalName)-2:]
		text = rx.ReplaceAllString(text, `${1}<num>`+directOrdinalStr+suffix+`${2}`)
	}

	for singleOrdinalName, singleOrdinalStr := range SINGLE_ORDINALS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + singleOrdinalName + `($|\W)`)
		suffix := singleOrdinalName[len(singleOrdinalName)-2:]
		text = rx.ReplaceAllString(text, `${1}<num>`+singleOrdinalStr+suffix+`${2}`)
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
		bigPrefixName := pair[0]
		bigPrefixValueStr := pair[1]
		rx := regexp.MustCompile(`(?i)(?:<num>)?(\d*) *` + bigPrefixName)
		matches := rx.FindStringSubmatch(text)

		if len(matches) == 2 {
			bigPrefixValue, _ := strconv.Atoi(bigPrefixValueStr)

			replacement := bigPrefixValueStr
			if len(matches[1]) > 0 {
				v1, _ := strconv.Atoi(matches[1])
				replacement = `<num>` + strconv.Itoa(v1*bigPrefixValue)
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
