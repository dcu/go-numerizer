package numerizer

var (
	DIRECT_NUMS = map[string]string{
		"eleven":    "11",
		"twelve":    "12",
		"thirteen":  "13",
		"fourteen":  "14",
		"fifteen":   "15",
		"sixteen":   "16",
		"seventeen": "17",
		"eighteen":  "18",
		"nineteen":  "19",
		"ninteen":   "19", // Common mis-spelling
		"zero":      "0",
		"ten":       "10",
		"\ba[\b^$]": "1", // doesn't make sense for an 'a' at the end to be a 1
	}

	SINGLE_NUMS = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	TEN_PREFIXES = map[string]string{
		"twenty":  "20",
		"thirty":  "30",
		"forty":   "40",
		"fourty":  "40", // Common misspelling
		"fifty":   "50",
		"sixty":   "60",
		"seventy": "70",
		"eighty":  "80",
		"ninety":  "90",
	}

	BIG_PREFIXES = [][]string{
		[]string{"hundred", "100"},
		[]string{"thousand", "1000"},
		[]string{"million", "1000000"},
		[]string{"billion", "1000000000"},
		[]string{"trillion", "1000000000000"},
	}

	FRACTIONS = map[string]string{
		"half":        "2",
		"third(s)?":   "3",
		"fourth(s)?":  "4",
		"quarter(s)?": "4",
		"fifth(s)?":   "5",
		"sixth(s)?":   "6",
		"seventh(s)?": "7",
		"eighth(s)?":  "8",
		"nineth(s)?":  "9",
	}

	SINGLE_ORDINALS = map[string]string{
		"first":   "1",
		"third":   "3",
		"fourth":  "4",
		"fifth":   "5",
		"sixth":   "6",
		"seventh": "7",
		"eighth":  "8",
		"ninth":   "9",
	}

	DIRECT_ORDINALS = map[string]string{
		"tenth":       "10",
		"eleventh":    "11",
		"twelfth":     "12",
		"thirteenth":  "13",
		"fourteenth":  "14",
		"fifteenth":   "15",
		"sixteenth":   "16",
		"seventeenth": "17",
		"eighteenth":  "18",
		"nineteenth":  "19",
		"twentieth":   "20",
		"thirtieth":   "30",
		"fourtieth":   "40",
		"fiftieth":    "50",
		"sixtieth":    "60",
		"seventieth":  "70",
		"eightieth":   "80",
		"ninetieth":   "90",
	}
)
