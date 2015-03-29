package numerizer

import (
	"strconv"
)

type NameValuePair struct {
	Name  string
	Value string
}

type NameValuePairs []*NameValuePair

func (pair *NameValuePair) ValueAsInt() int {
	value, err := strconv.Atoi(pair.Value)
	if err != nil {
		return 0.0
	}

	return value
}
