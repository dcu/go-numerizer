package numerizer

type NameValuePair struct {
	Name  string
	Value string
}

type NameValuePairs []*NameValuePair

func (pair *NameValuePair) ValueAsInt() int {
	return stringToInt(pair.Value)
}
