package iban_test

import (
	"testing"

	"github.com/ritchieflick/iban"
)

var ibanTestNumbers = []struct {
	number string
}{
	{"LU28 0019 4006 4475 0000"},
	{"LU12 3456 7890 1234 5678"},
}

func TestIBAN(t *testing.T) {
	for _, ibanTestNumber := range ibanTestNumbers {
		result, err := iban.NewIBAN(ibanTestNumber.number)
		if err != nil || result == (iban.IBAN{}) {
			t.Error("No object was created!")
			t.Log(err)
		}
	}
}
