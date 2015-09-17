// Package iban_check contains utility functions for working with IBAN account numbers.
package iban

import (
	"fmt"
	"strconv"
	"strings"
)

//****************** IsCorrectIban **********************************************//
func IsCorrectIban(iban string) (bool, error) {
	iban_valid := false

	if len(iban) > 14 {

		iban = strings.Replace(iban, " ", "", -1)
		iban = rearrange(iban)
		iban = convert_char_to_number(iban)

		if calculate_modulo(iban) == 1 {
			iban_valid = true
		}

	} else {
		return false, fmt.Errorf("IBAN: incorect IBAN string passed <%s>", iban)
	}
	return iban_valid, nil
}

//****************** GetIbanChecksum **********************************************//
func GetIbanChecksum(iban string) (int, error) {
	iban_checksum := 0

	if len(iban) > 14 {

		iban = strings.Replace(iban, " ", "", -1)
		iban = rearrange(iban)
		iban = convert_char_to_number(iban)

		iban_checksum = 98 - calculate_modulo(iban)

	} else {
		return -1, fmt.Errorf("IBAN: incorect IBAN string passed <%s>", iban)
	}
	return iban_checksum, nil
}

func convert_char_to_number(value string) string {
	var temp []byte = []byte(value)
	var result string

	for _, item := range temp {
		if !((item < "A"[0]) || (item > "Z"[0])) {
			item = byte(item - 55)
			result += strconv.Itoa(int(item))
		} else {
			result += string(item)
		}
	}
	return result
}

func rearrange(iban string) string {
	country := iban[0:2]
	checksum := iban[2:4]
	body := iban[4:len(iban)]
	return body + country + checksum
}

func calculate_modulo(iban string) int {
	exit := false
	tempIBAN := ""
	rest := 0
	for !exit {
		if len(iban) > 9 {
			tempIBAN = iban[0:9]

			value, err := strconv.Atoi(tempIBAN)
			if err != nil {
				exit = true
				return -1
			}

			rest = (value % 97)
			iban = strconv.Itoa(rest) + iban[9:len(iban)]
		} else {
			tempIBAN = iban

			value, err := strconv.Atoi(tempIBAN)
			if err != nil {
				exit = true
				return -1
			}

			rest = (value % 97)
			exit = true
		}
	}
	return rest
}
