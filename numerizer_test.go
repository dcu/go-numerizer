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
