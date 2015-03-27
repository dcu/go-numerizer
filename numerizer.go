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
		text = rx.ReplaceAllString(text, `${1}<num>`+value)
	}

	for key, value := range SINGLE_NUMS {
		rx := regexp.MustCompile(`(?i)(^|\W)` + key + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+value)
	}

	for tenPrefixName, tenPrefixStr := range TEN_PREFIXES {
		for singleNumName, singleNumStr := range SINGLE_NUMS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + singleNumName + `($|\W)`)

			tenPrefixNum, _ := strconv.Atoi(tenPrefixStr)
			singleNumNum, _ := strconv.Atoi(singleNumStr)

			num := strconv.Itoa(tenPrefixNum + singleNumNum)

			text = rx.ReplaceAllString(text, `${1}<num>`+num)
		}

		for singleOrdinalName, singleOrdinalStr := range SINGLE_ORDINALS {
			rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixName + `(\s)?` + singleOrdinalName + `($|\W)`)

			tenPrefixNum, _ := strconv.Atoi(tenPrefixStr)
			singleOrdinalNum, _ := strconv.Atoi(singleOrdinalStr)

			num := strconv.Itoa(tenPrefixNum + singleOrdinalNum)
			suffix := singleOrdinalName[len(singleOrdinalName)-2:]

			text = rx.ReplaceAllString(text, `${1}<num>`+num+suffix)
		}

		rx := regexp.MustCompile(`(?i)(^|\W)` + tenPrefixName + `($|\W)`)
		text = rx.ReplaceAllString(text, `${1}<num>`+tenPrefixStr)
	}

	text = strings.Replace(text, "<num>", "", -1)

	return text, err
}

func andition(string) {

}
