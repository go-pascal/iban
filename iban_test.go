package iban

import (
	"testing"
)

var validIBANTestNumbers = []struct {
	number string
}{
	{"LU28 0019 4006 4475 0000"},
}

var invalidIBANTestNumbers = []struct {
	number string
}{
	{"LU12 3456 7890 1234 5678"},
}

func TestValidIBAN(t *testing.T) {
	for _, ibanTestNumber := range validIBANTestNumbers {
		result, err := NewIBAN(ibanTestNumber.number)
		if err != nil || result == (IBAN{}) {
			t.Error("No object was created!")
			t.Log(err)
		}
	}
}

func TestInvalidIBAN(t *testing.T) {
	for _, ibanTestNumber := range invalidIBANTestNumbers {
		_, err := NewIBAN(ibanTestNumber.number)
		if err == nil {
			t.Error("No error was thrown for an invalid IBAN number!")
		}
	}
}
