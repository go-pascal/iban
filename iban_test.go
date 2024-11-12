package iban

import (
	"encoding/json"
	"os"
	"testing"
)

var validIBANTestNumbers = []struct {
	number string
}{
	{"EG210003700067100239218937900"},
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

// TestIsCorrectIban reads a list of IBANs from JSON and verifies each one with detailed logging.
func TestIsCorrectIban(t *testing.T) {
	t.Log("Starting TestIsCorrectIban")

	// Attempt to open the JSON file
	t.Log("Attempting to open './data/iban.json'")
	file, err := os.Open("./data/iban.json")
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	t.Log("Successfully opened './data/iban.json'")

	// Decode JSON data
	t.Log("Decoding JSON data from file")
	var iList IBANList
	if err := json.NewDecoder(file).Decode(&iList); err != nil {
		t.Fatalf("Error unmarshalling data into the IBAN list: %v", err)
	}
	t.Log("Successfully decoded JSON data")

	// Test each IBAN entry
	for i, message := range iList.IBANs {
		//t.Logf("Testing IBAN: %s (Country: %s, Code: %s)", message.IBAN, message.Country, message.Code)
		ok, _, _ := IsCorrectIban(message.IBAN, false)
		if !ok {
			t.Errorf("Test case %d: Expected valid IBAN for %+v, but got invalid.", i+1, message)
		}
	}

	t.Log("Finished TestIsCorrectIban")
}
