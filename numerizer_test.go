package numerizer

import (
	"testing"
)

func Test_DirectNums(t *testing.T) {
	number, err := Numerize("eleven")

	if err != nil {
		t.Error("Number not parsed, error: " + err.Error())
	}

	if number != "11" {
		t.Error("Number not parsed")
	}
}

func Test_TenPrefix(t *testing.T) {
	number, err := Numerize("thirty")

	if err != nil {
		t.Error("Number not parsed, error: " + err.Error())
	}

	if number != "30" {
		t.Error("Number not parsed")
	}
}

func Test_Fraction_1(t *testing.T) {
	number, err := Numerize("a fifth")

	if err != nil {
		t.Error("Number not parsed, error: " + err.Error())
	}

	if number != "1/5" {
		t.Error("Number not parsed")
	}
}

func Test_Fractions_2(t *testing.T) {
	number, err := Numerize("two fifths")

	if err != nil {
		t.Error("Number not parsed, error: " + err.Error())
	}

	if number != "2/5" {
		t.Error("Number not parsed")
	}
}

func Test_DirectOrdinal(t *testing.T) {
	number, err := Numerize("ninth")

	if err != nil {
		t.Error("Number not parsed, error: " + err.Error())
	}

	if number != "9th" {
		t.Error("Number not parsed")
	}
}

func Test_SingleOrdinal(t *testing.T) {
	number, err := Numerize("third")

	if err != nil {
		t.Error("Number not parsed, error: " + err.Error())
	}

	if number != "3rd" {
		t.Error("Number not parsed")
	}
}

func Test_CleanFraction(t *testing.T) {
	number, err := Numerize("one and two fifths")

	if err != nil {
		t.Error("Number not parsed, error: " + err.Error())
	}

	if number != "1.40" {
		t.Error("Number not parsed")
	}
}
