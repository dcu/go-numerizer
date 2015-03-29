package numerizer

import (
	"testing"
)

var (
	Examples = map[string]string{
		"one":            "1",
		"two and a half": "2.5",
		"five":           "5",
		"ten":            "10",
		"eleven":         "11",
		"twelve":         "12",
		"thirteen":       "13",
		"fourteen":       "14",
		"fifteen":        "15",
		"sixteen":        "16",
		"seventeen":      "17",
		"eighteen":       "18",
		"nineteen":       "19",
		"twenty":         "20",
		"twenty seven":   "27",
		"thirty-one":     "31",
		"thirty-seven":   "37",
		"forty one":      "41",
		"fourty two":     "42",
		"fifty nine":     "59",
		//"a hundred":                                         "100",
		"one hundred":                                       "100",
		"one hundred and fifty":                             "150",
		"two-hundred":                                       "200",
		"5 hundred":                                         "500",
		"nine hundred and ninety nine":                      "999",
		"one thousand":                                      "1000",
		"twelve hundred":                                    "1200",
		"one thousand two hundred":                          "1200",
		"seventeen thousand":                                "17000",
		"twentyone-thousand-four-hundred-and-seventy-three": "21473",
		"seventy four thousand and two":                     "74002",
		"ninety nine thousand nine hundred ninety nine":     "99999",
		"100 thousand":                                      "100000",
		"two hundred fifty thousand":                        "250000",
		"one million":                                       "1000000",
		"one million two hundred fifty thousand and seven":  "1250007",
		"one billion":                                       "1000000000",
		"one billion and one":                               "1000000001",
	}

	Ordinals = map[string]string{
		"first":               "1st",
		"second":              "second",
		"third":               "3rd",
		"fifth":               "5th",
		"seventh":             "7th",
		"eighth":              "8th",
		"tenth":               "10th",
		"eleventh":            "11th",
		"twelfth":             "12th",
		"thirteenth":          "13th",
		"sixteenth":           "16th",
		"twentieth":           "20th",
		"twenty-third":        "23rd",
		"thirtieth":           "30th",
		"thirty-first":        "31st",
		"fourtieth":           "40th",
		"fourty ninth":        "49th",
		"fiftieth":            "50th",
		"sixtieth":            "60th",
		"seventieth":          "70th",
		"eightieth":           "80th",
		"ninetieth":           "90th",
		"hundredth":           "100th",
		"thousandth":          "1000th",
		"millionth":           "1000000th",
		"billionth":           "1000000000th",
		"trillionth":          "1000000000000th",
		"first day month two": "1st day month 2",
	}
)

func Test_CombinedDoubleDigits(t *testing.T) {
	number := Numerize("thirtyseven")

	if number != "37" {
		t.Error("Number not parsed")
	}

}

//func Test_A_Million(t *testing.T) {
//number, err := Numerize("a million")

//if err != nil {
//t.Error("Number not parsed, error: " + err.Error())
//}

//if number != "1000000" {
//t.Error("Number not parsed")
//}
//}

func Test_DirectNums(t *testing.T) {
	number := Numerize("eleven")

	if number != "11" {
		t.Error("Number not parsed")
	}
}

func Test_TenPrefix(t *testing.T) {
	number := Numerize("thirty")

	if number != "30" {
		t.Error("Number not parsed")
	}
}

func Test_Fraction_1(t *testing.T) {
	number := Numerize("a fifth")

	if number != "1/5" {
		t.Error("Number not parsed")
	}
}

func Test_Fractions_2(t *testing.T) {
	number := Numerize("two fifths")

	if number != "2/5" {
		t.Error("Number not parsed")
	}
}

func Test_DirectOrdinal(t *testing.T) {
	number := Numerize("ninth")

	if number != "9th" {
		t.Error("Number not parsed")
	}
}

func Test_SingleOrdinal(t *testing.T) {
	number := Numerize("third")

	if number != "3rd" {
		t.Error("Number not parsed")
	}
}

func Test_CleanFraction(t *testing.T) {
	number := Numerize("one and two fifths")

	if number != "1.4" {
		t.Error("Number not parsed")
	}
}

func Test_BigPrefix(t *testing.T) {
	number := Numerize("two trillion")

	if number != "2000000000000" {
		t.Error("Number not parsed")
	}
}

func Test_Numbers(t *testing.T) {
	for text, numberStr := range Examples {
		result := Numerize(text)

		if result != numberStr {
			t.Errorf("Failed test. %s:  %s != %s", text, numberStr, result)
		}
	}
}

func Test_Ordinals(t *testing.T) {
	for text, numberStr := range Ordinals {
		result := Numerize(text)

		if result != numberStr {
			t.Errorf("Failed test. %s:  %s != %s", text, numberStr, result)
		}
	}
}

func Test_Compatability(t *testing.T) {
	if "1/2" != Numerize("1/2") {
		t.Error("Number not parsed correctly.")
	}

	if "05/06" != Numerize("05/06") {
		t.Error("Number not parsed correctly.")
	}

	if "3.5 hours" != Numerize("three and a half hours") {
		println(Numerize("three and a half hours"))
		t.Error("Number not parsed correctly.")
	}
}

func Test_MultipleSlashesShouldNotBeEvaluated(t *testing.T) {
	if "11/02/2007" != Numerize("11/02/2007") {
		t.Error("Number not parsed correctly.")
	}
}

func Test_Edges(t *testing.T) {
	if "27 Oct 2006 7:30am" != Numerize("27 Oct 2006 7:30am") {
		t.Error("Number not parsed correctly.")
	}
}

func Test_WordWithANumber(t *testing.T) {
	if "pennyweight" != Numerize("pennyweight") {
		t.Error("Number not parsed correctly.")
	}
}

func Test_FractionalAddition(t *testing.T) {
	if "1.25" != Numerize("one and a quarter") {
		t.Error("Number not parsed correctly.")
	}
	if "2.375" != Numerize("two and three eighths") {
		t.Error("Number not parsed correctly.")
	}
	if "3.5 hours" != Numerize("three and a half hours") {
		t.Error("Number not parsed correctly.")
	}
}

func Test_CombinedDoubleDigets(t *testing.T) {
	if "21" != Numerize("twentyone") {
		t.Error("Number not parsed correctly.")
	}
	if "37" != Numerize("thirtyseven") {
		t.Error("Number not parsed correctly.")
	}
}

func Test_FractionsInWords(t *testing.T) {
	if "1/4" != Numerize("1 quarter") {
		t.Error("Number not parsed correctly.")
	}
	if "1/4" != Numerize("one quarter") {
		t.Error("Number not parsed correctly.")
	}
	if "1/4" != Numerize("a quarter") {
		t.Error("Number not parsed correctly.")
	}
	if "1/8" != Numerize("one eighth") {
		t.Error("Number not parsed correctly.")
	}

	if "3/4" != Numerize("three quarters") {
		t.Error("Number not parsed correctly.")
	}
	if "2/4" != Numerize("two fourths") {
		t.Error("Number not parsed correctly.")
	}
	if "3/8" != Numerize("three eighths") {
		t.Error("Number not parsed correctly.")
	}
}
