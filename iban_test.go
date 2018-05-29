package iban

import (
	"encoding/json"
	"io/ioutil"
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

type IBANMessage struct {
	Country string `json:"country"`
	Code    string `json:"code"`
	IBAN    string `json:"iban"`
}

type IBANList struct {
	IBANs []IBANMessage `json:"ibans"`
}

func TestIsCorrectIban(t *testing.T) {

	data, err := ioutil.ReadFile("./data/iban.json")
	if err != nil {
		t.Error("error reading file", err)
		t.FailNow()
	}
	var iList IBANList
	err = json.Unmarshal(data, &iList)
	if err != nil {
		t.Error("error unmarshall data into the IBAN list", err)
		t.FailNow()
	}

	for k, message := range iList.IBANs {
		ok, _, _ := IsCorrectIban(message.IBAN, false)
		if !ok {
			t.Error("for test waiting for true got false", k, ":", message)
		}
	}
}
