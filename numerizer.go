package numerizer

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	HyphenatedWordsRx = regexp.MustCompile(` +|([^\d])-([^\d])`)
	AndRx             = regexp.MustCompile(`(?i)<num>(\d+)( | and )<num>(\d+)([^\w]|$)`)
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
			rx := regexp.MustCompile(`(?i)(^|\W)` + singleNumName + `($|\W)`)

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
	text = strings.Replace(text, "<num>", "", -1)

	return text, err
}

func andition(string) {

}
