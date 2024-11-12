// Package iban contains utility functions for working with IBAN account numbers.
package iban

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ErrInvalidIBAN is returned when an invalid IBAN number was received
var ErrInvalidIBAN = errors.New("invalid IBAN number received")

// IBAN represents an IBAN number, split up into its different parts.
type IBAN struct {
	Number      string // The full IBAN number
	CountryCode string // The country code of the IBAN
	Checksum    string // The checksum of the IBAN
	BBAN        string // The Basic Bank Account Number (BBAN) of the IBAN
	BankCode    string // The bank code extracted from the IBAN
}

// NewIBAN creates a new instance of IBAN and checks if the IBAN number is valid.
// If the IBAN is valid, it returns the IBAN struct with its different parts filled in.
func NewIBAN(ibanNumber string) (IBAN, error) {
	isIBANValid, formattedIBANNumber, err := IsCorrectIban(ibanNumber, false)
	if err != nil || !isIBANValid {
		return IBAN{}, fmt.Errorf("%w: %v", ErrInvalidIBAN, err)
	}

	countryCode, checksum, bban := splitIbanUp(formattedIBANNumber)
	bankCode, _ := getBankCode(countryCode, formattedIBANNumber)

	return IBAN{
		Number:      formattedIBANNumber,
		CountryCode: countryCode,
		Checksum:    checksum,
		BBAN:        bban,
		BankCode:    bankCode,
	}, nil
}

// getBankCode extracts the bank code from the IBAN based on the country configuration.
func getBankCode(countryCode, ibanNumber string) (string, error) {
	if code, exists := countryList[countryCode]; exists {
		firstIndex := strings.Index(code.ibanFields, "b")
		lastIndex := strings.LastIndex(code.ibanFields, "b")
		if firstIndex != -1 && lastIndex != -1 {
			return ibanNumber[firstIndex : lastIndex+1], nil
		}
	}
	return "", errors.New("no bank code found")
}

type ibanCountry struct {
	country           string // The country name
	chars             int    // The expected length of the IBAN for this country
	bbanFormat        string // The format of the BBAN part of the IBAN
	code              string // The country code
	ibanFields        string // The fields of the IBAN (e.g., bank code, branch code)
	comment           string // Additional comments about the IBAN format
	standardTreatment bool   // Indicates if the country follows the standard treatment
}

// IsCorrectIban checks if the given IBAN number corresponds to the rules of a valid IBAN number.
// It also returns a properly formatted IBAN number string.
func IsCorrectIban(iban string, debug bool) (bool, string, error) {
	obscureIban := func(iban string) string {
		if len(iban) <= 6 {
			return iban
		}
		return iban[:2] + strings.Repeat("*", len(iban)-6) + iban[len(iban)-4:]
	}

	// Clean up and standardize the IBAN string
	iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))
	if len(iban) < 15 {
		log.Printf("Incorrect IBAN string passed: %s", obscureIban(iban))
		return false, "", fmt.Errorf("IBAN: incorrect IBAN string passed <%s>", obscureIban(iban))
	}

	// Split the IBAN into its parts
	countryCode, checksum, bban := splitIbanUp(iban)
	ibanConfig, exists := countryList[countryCode]
	if !exists {
		log.Printf("Invalid IBAN country: IBAN %s has country code not in the list (%s)", obscureIban(iban), countryCode)
		return false, "", fmt.Errorf("IBAN: country <%s> is not in the list", countryCode)
	}

	// Check if the length matches the expected length for the country
	if ibanConfig.chars != len(iban) {
		log.Printf("Invalid IBAN length: IBAN %s length (%d) does not match configuration length (%d)", obscureIban(iban), len(iban), ibanConfig.chars)
		return false, "", fmt.Errorf("IBAN: length (%d) does not match configuration length (%d)", len(iban), ibanConfig.chars)
	}

	// Rearrange the IBAN for validation and convert characters to numbers
	rearrangedIban := rearrangeIBAN(countryCode, checksum, bban)
	convertedIban := convertCharToNumber(rearrangedIban)
	if calculateModulo(convertedIban) != 1 {
		return false, "", nil
	}

	return true, splitTo4(iban), nil
}

// splitIbanUp splits the IBAN into its country code, checksum, and BBAN parts.
func splitIbanUp(iban string) (string, string, string) {
	return iban[:2], iban[2:4], iban[4:]
}

// splitTo4 splits the IBAN into groups of four characters for readability.
func splitTo4(value string) string {
	var builder strings.Builder
	for i := 0; i < len(value); i += 4 {
		if i+4 < len(value) {
			builder.WriteString(value[i:i+4] + " ")
		} else {
			builder.WriteString(value[i:])
		}
	}
	return builder.String()
}

// GetIbanChecksum returns the checksum of the given IBAN number.
func GetIbanChecksum(iban string) (int, error) {
	// Clean up and standardize the IBAN string
	iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))
	if len(iban) < 15 {
		log.Printf("Incorrect IBAN string passed for checksum calculation: %s", obscureIban(iban))
		return -1, fmt.Errorf("IBAN: incorrect IBAN string passed <%s>", obscureIban(iban))
	}

	// Rearrange the IBAN for checksum calculation
	countryCode, _, bban := splitIbanUp(iban)
	rearrangedIban := rearrangeIBAN(countryCode, "00", bban)
	convertedIban := convertCharToNumber(rearrangedIban)
	return 98 - calculateModulo(convertedIban), nil
}

// convertCharToNumber converts alphabetic characters in the IBAN to their corresponding numeric values.
func convertCharToNumber(value string) string {
	var builder strings.Builder
	for _, item := range value {
		if item >= 'A' && item <= 'Z' {
			builder.WriteString(strconv.Itoa(int(item - 'A' + 10)))
		} else {
			builder.WriteRune(item)
		}
	}
	return builder.String()
}

// rearrangeIBAN rearranges the IBAN by moving the country code and checksum to the end.
func rearrangeIBAN(countryCode, checksum, bban string) string {
	return bban + countryCode + checksum
}

// calculateModulo calculates the modulo 97 of the given IBAN number.
func calculateModulo(iban string) int {
	for len(iban) > 9 {
		tempIBAN := iban[:9]
		value, err := strconv.Atoi(tempIBAN)
		if err != nil {
			return -1
		}
		rest := value % 97
		iban = strconv.Itoa(rest) + iban[9:]
	}

	value, err := strconv.Atoi(iban)
	if err != nil {
		return -1
	}

	return value % 97
}

// obscureIban obscures the middle part of an IBAN for logging purposes.
func obscureIban(iban string) string {
	if len(iban) <= 6 {
		return iban
	}
	return iban[:2] + strings.Repeat("*", len(iban)-6) + iban[len(iban)-4:]
}
