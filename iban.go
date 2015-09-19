// Package iban_check contains utility functions for working with IBAN account numbers.
package iban

import (
	"fmt"
	"strconv"
	"strings"
)

type iban_country struct {
	country            string
	chars              int
	bban_format        string
	code               string
	iban_fields        string
	comment            string
	standard_treatment bool
}

func (s iban_country) WriteCountryInfo() {
	fmt.Println("Country      : " + s.country)
	fmt.Println("Country code : " + s.code)
	fmt.Println("Length       : " + strconv.Itoa(s.chars))
	fmt.Println("IBAN fields  : " + s.iban_fields)
	fmt.Println("BBAN format  : " + s.bban_format)
	fmt.Println("Comment      : " + s.comment)
}

var (
	passed_code     string
	passed_chars    int
	passed_checksum string
	passeds_bban    string

	country_list = map[string]iban_country{
		"AL": iban_country{country: "Albania", chars: 28, bban_format: "8n, 16c", code: "AL", iban_fields: "ALkk bbbs sssx cccc cccc cccc cccc", comment: "b = National bank code s = Branch code x = National check digit c = Account number", standard_treatment: true},
		"AD": iban_country{country: "Andorra", chars: 24, bban_format: "8n,12c", code: "AD", iban_fields: "ADkk bbbb ssss cccc cccc cccc", comment: "b = National bank code s = Branch code c = Account number", standard_treatment: true},
		"AT": iban_country{country: "Austria", chars: 20, bban_format: "16n", code: "AT", iban_fields: "ATkk bbbb bccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"AZ": iban_country{country: "Azerbaijan", chars: 28, bban_format: "4c,20n", code: "AZ", iban_fields: "AZkk bbbb cccc cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"BH": iban_country{country: "Bahrain", chars: 22, bban_format: "4a,14c", code: "BH", iban_fields: "BHkk bbbb cccc cccc cccc cc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"BE": iban_country{country: "Belgium", chars: 16, bban_format: "12n", code: "BE", iban_fields: "BEkk bbbc cccc ccxx", comment: "b = National bank code c = Account number x = National check digits", standard_treatment: true},
		"BA": iban_country{country: "Bosnia and Herzegovina", chars: 20, bban_format: "16n", code: "BA", iban_fields: "BAkk bbbs sscc cccc ccxx", comment: "k = IBAN check digits (always 39) b = National bank code s = Branch code c = Account number x = National check digits", standard_treatment: true},
		"BR": iban_country{country: "Brazil", chars: 29, bban_format: "23n, 1a, 1c", code: "BR", iban_fields: "BRkk bbbb bbbb ssss sccc cccc ccct n", comment: "k = IBAN check digits (Calculated by MOD 97-10) b = National bank code s = Branch code c = Account Number t = Account type (Cheque account, Savings account etc.) n = Owner account number (1, 2 etc.)[31]", standard_treatment: true},
		"BG": iban_country{country: "Bulgaria", chars: 22, bban_format: "4a,6n,8c", code: "BG", iban_fields: "BGkk bbbb ssss ddcc cccc cc", comment: "b = BIC bank code s = Branch (BAE) number d = Account type c = Account number", standard_treatment: true},
		"CR": iban_country{country: "Costa Rica", chars: 21, bban_format: "17n", code: "CR", iban_fields: "CRkk bbbc cccc cccc cccc c", comment: "b = bank code c = Account number", standard_treatment: true},
		"HR": iban_country{country: "Croatia", chars: 21, bban_format: "17n", code: "HR", iban_fields: "HRkk bbbb bbbc cccc cccc c", comment: "b = Bank code c = Account number", standard_treatment: true},
		"CY": iban_country{country: "Cyprus", chars: 28, bban_format: "8n,16c", code: "CY", iban_fields: "CYkk bbbs ssss cccc cccc cccc cccc", comment: "b = National bank code s = Branch code c = Account number", standard_treatment: true},
		"CZ": iban_country{country: "Czech Republic", chars: 24, bban_format: "20n", code: "CZ", iban_fields: "CZkk bbbb ssss sscc cccc cccc", comment: "b = National bank code s = Account number prefix c = Account number", standard_treatment: true},
		"DK": iban_country{country: "Denmark", chars: 18, bban_format: "14n", code: "DK", iban_fields: "DKkk bbbb cccc cccc cc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"DO": iban_country{country: "Dominican Republic", chars: 28, bban_format: "4a,20n", code: "DO", iban_fields: "DOkk bbbb cccc cccc cccc cccc cccc", comment: "b = Bank identifier c = Account number", standard_treatment: true},
		"TL": iban_country{country: "East Timor", chars: 23, bban_format: "19n", code: "TL", iban_fields: "TLkk bbbc cccc cccc cccc cxx", comment: "k = IBAN check digits (always = '38') b = Bank identifier c = Account number x = National check digit", standard_treatment: true},
		"EE": iban_country{country: "Estonia", chars: 20, bban_format: "16n", code: "EE", iban_fields: "EEkk bbss cccc cccc cccx", comment: "b = National bank code s = Branch code c = Account number x = National check digit", standard_treatment: true},
		"FO": iban_country{country: "Faroe Islands", chars: 18, bban_format: "14n", code: "FO", iban_fields: "FOkk bbbb cccc cccc cx", comment: "b = National bank code c = Account number x = National check digit", standard_treatment: true},
		"FI": iban_country{country: "Finland", chars: 18, bban_format: "14n", code: "FI", iban_fields: "FIkk bbbb bbcc cccc cx", comment: "b = Bank and branch code c = Account number x = National check digit", standard_treatment: true},
		"FR": iban_country{country: "France", chars: 27, bban_format: "10n,11c,2n", code: "FR", iban_fields: "FRkk bbbb bggg ggcc cccc cccc cxx", comment: "b = National bank code g = Branch code (fr:code guichet) c = Account number x = National check digits (fr:clé RIB)", standard_treatment: true},
		"GE": iban_country{country: "Georgia", chars: 22, bban_format: "2c,16n", code: "GE", iban_fields: "GEkk bbcc cccc cccc cccc cc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"DE": iban_country{country: "Germany", chars: 22, bban_format: "18n", code: "DE", iban_fields: "DEkk bbbb bbbb cccc cccc cc", comment: "b = Bank and branch identifier (de:Bankleitzahl or BLZ) c = Account number", standard_treatment: true},
		"GI": iban_country{country: "Gibraltar", chars: 23, bban_format: "4a,15c", code: "GI", iban_fields: "GIkk bbbb cccc cccc cccc ccc", comment: "b = BIC bank code c = Account number", standard_treatment: true},
		"GR": iban_country{country: "Greece", chars: 27, bban_format: "7n,16c", code: "GR", iban_fields: "GRkk bbbs sssc cccc cccc cccc ccc", comment: "b = National bank code s = Branch code c = Account number", standard_treatment: true},
		"GL": iban_country{country: "Greenland", chars: 18, bban_format: "14n", code: "GL", iban_fields: "GLkk bbbb cccc cccc cc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"GT": iban_country{country: "Guatemala", chars: 28, bban_format: "4c,20c", code: "GT", iban_fields: "GTkk bbbb mmtt cccc cccc cccc cccc", comment: "b = National bank code c = Account number m = Currency t = Account type ", standard_treatment: true},
		"HU": iban_country{country: "Hungary", chars: 28, bban_format: "24n", code: "HU", iban_fields: "HUkk bbbs sssk cccc cccc cccc cccx", comment: "b = National bank code s = Branch code c = Account number x = National check digit", standard_treatment: true},
		"IS": iban_country{country: "Iceland", chars: 26, bban_format: "22n", code: "IS", iban_fields: "ISkk bbbb sscc cccc iiii iiii ii", comment: "b = National bank code s = Branch code c = Account number i = holder's kennitala (national identification number).", standard_treatment: true},
		"IE": iban_country{country: "Ireland", chars: 22, bban_format: "4c,14n", code: "IE", iban_fields: "IEkk aaaa bbbb bbcc cccc cc", comment: "a = BIC bank code b = Bank/branch code (sort code) c = Account number", standard_treatment: true},
		"IK": iban_country{country: "Israel", chars: 23, bban_format: "19n", code: "IK", iban_fields: "ILkk bbbn nncc cccc cccc ccc", comment: "b = National bank code n = Branch number c = Account number 13 digits (padded with zeros)", standard_treatment: true},
		"IT": iban_country{country: "Italy", chars: 27, bban_format: "1a,10n,12c", code: "IT", iban_fields: "ITkk xaaa aabb bbbc cccc cccc ccc", comment: "x = Check char (CIN) a = National bank code (it:Associazione bancaria italiana or Codice ABI ) b = Branch code (it:Coordinate bancarie or CAB – Codice d'Avviamento Bancario) c = Account number", standard_treatment: true},
		"JO": iban_country{country: "Jordan", chars: 30, bban_format: "4a, 22n", code: "JO", iban_fields: "JOkk bbbb nnnn cccc cccc cccc cccc cc", comment: "b = National bank code n = Branch code c = Account number ", standard_treatment: true},
		"KZ": iban_country{country: "Kazakhstan", chars: 20, bban_format: "3n,13c", code: "KZ", iban_fields: "KZkk bbbc cccc cccc cccc", comment: "b = National bank code c = Account number ", standard_treatment: true},
		"XK": iban_country{country: "Kosovo", chars: 20, bban_format: "4n,10n,2n", code: "XK", iban_fields: "XKkk bbbb cccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"KW": iban_country{country: "Kuwait", chars: 30, bban_format: "4a, 22c", code: "KW", iban_fields: "KWkk bbbb cccc cccc cccc cccc cccc cc", comment: "b = National bank code c = Account number.", standard_treatment: true},
		"LV": iban_country{country: "Latvia", chars: 21, bban_format: "4a,13c", code: "LV", iban_fields: "LVkk bbbb cccc cccc cccc c", comment: "b = BIC Bank code c = Account number", standard_treatment: true},
		"LB": iban_country{country: "Lebanon", chars: 28, bban_format: "4n,20c", code: "LB", iban_fields: "LBkk bbbb cccc cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"LI": iban_country{country: "Liechtenstein", chars: 21, bban_format: "5n,12c", code: "LI", iban_fields: "LIkk bbbb bccc cccc cccc c", comment: "b = National bank code c = Account number", standard_treatment: true},
		"LT": iban_country{country: "Lithuania", chars: 20, bban_format: "16n", code: "LT", iban_fields: "LTkk bbbb bccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"LU": iban_country{country: "Luxembourg", chars: 20, bban_format: "3n,13c", code: "LU", iban_fields: "LUkk bbbc cccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"MK": iban_country{country: "Macedonia", chars: 19, bban_format: "3n,10c,2n", code: "MK", iban_fields: "MKkk bbbc cccc cccc cxx", comment: "k = IBAN check digits (always = '07') b = National bank code c = Account number x = National check digits", standard_treatment: true},
		"MT": iban_country{country: "Malta", chars: 31, bban_format: "4a,5n,18c", code: "MT", iban_fields: "MTkk bbbb ssss sccc cccc cccc cccc ccc", comment: "b = BIC bank code s = Branch code c = Account number", standard_treatment: true},
		"MR": iban_country{country: "Mauritania", chars: 27, bban_format: "23n", code: "MR", iban_fields: "MRkk bbbb bsss sscc cccc cccc cxx", comment: "k = IBAN check digits (always 13) b = National bank code s = Branch code (fr:code guichet) c = Account number x = National check digits (fr:clé RIB)", standard_treatment: true},
		"MU": iban_country{country: "Mauritius", chars: 30, bban_format: "4a,19n,3a", code: "MU", iban_fields: "MUkk bbbb bbss cccc cccc cccc 000d dd", comment: "b = National bank code s = Branch identifier c = Account number 0 = Zeroes d = Currency Symbol ", standard_treatment: true},
		"MC": iban_country{country: "Monaco", chars: 27, bban_format: "10n,11c,2n", code: "MC", iban_fields: "MCkk bbbb bsss sscc cccc cccc cxx", comment: "b = National bank code s = Branch code (fr:code guichet) c = Account number x = National check digits (fr:clé RIB). ", standard_treatment: true},
		"MD": iban_country{country: "Moldova", chars: 24, bban_format: "2c,18c", code: "MD", iban_fields: "MDkk bbcc cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"ME": iban_country{country: "Montenegro", chars: 22, bban_format: "18n", code: "ME", iban_fields: "MEkk bbbc cccc cccc cccc xx", comment: "k = IBAN check digits (always = '25') b = Bank code c = Account number x = National check digits", standard_treatment: true},
		"NL": iban_country{country: "Netherlands", chars: 18, bban_format: "4a,10n", code: "NL", iban_fields: "NLkk bbbb cccc cccc cc", comment: "b = BIC Bank code c = Account number", standard_treatment: true},
		"NO": iban_country{country: "Norway", chars: 15, bban_format: "11n", code: "NO", iban_fields: "NOkk bbbb cccc ccx", comment: "b = National bank code c = Account number x = Modulo-11 national check digit", standard_treatment: true},
		"PK": iban_country{country: "Pakistan", chars: 24, bban_format: "4c,16n", code: "PK", iban_fields: "PKkk bbbb cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
		"PS": iban_country{country: "Palestinian territories", chars: 29, bban_format: "4c,21n", code: "PS", iban_fields: "PSkk bbbb xxxx xxxx xccc cccc cccc c", comment: "b = National bank code c = Account number x = Not specified", standard_treatment: true},
		"PL": iban_country{country: "Poland", chars: 28, bban_format: "24n", code: "PL", iban_fields: "PLkk bbbs sssx cccc cccc cccc cccc", comment: "b = National bank code s = Branch code x = National check digit c = Account number, ", standard_treatment: true},
		"PT": iban_country{country: "Portugal", chars: 25, bban_format: "21n", code: "PT", iban_fields: "PTkk bbbb ssss cccc cccc cccx x", comment: "k = IBAN check digits (always = '50') b = National bank code s = Branch code C = Account number x = National check digit", standard_treatment: true},
		"QA": iban_country{country: "Qatar", chars: 29, bban_format: "4a, 21c", code: "QA", iban_fields: "QAkk bbbb cccc cccc cccc cccc cccc c", comment: "b = National bank code c = Account number[34]", standard_treatment: true},
		"RO": iban_country{country: "Romania", chars: 24, bban_format: "4a,16c", code: "RO", iban_fields: "ROkk bbbb cccc cccc cccc cccc", comment: "b = BIC Bank code c = Branch code and account number (bank-specific format) ", standard_treatment: true},
		"SM": iban_country{country: "San Marino", chars: 27, bban_format: "1a,10n,12c", code: "SM", iban_fields: "SMkk xaaa aabb bbbc cccc cccc ccc", comment: "x = Check char (it:CIN) a = National bank code (it:Associazione bancaria italiana or Codice ABI) b = Branch code (it:Coordinate bancarie or CAB – Codice d'Avviamento Bancario) c = Account number", standard_treatment: true},
		"SA": iban_country{country: "Saudi Arabia", chars: 24, bban_format: "2n,18c", code: "SA", iban_fields: "SAkk bbcc cccc cccc cccc cccc", comment: "b = National bank code c = Account number preceded by zeros, if required", standard_treatment: true},
		"RS": iban_country{country: "Serbia", chars: 22, bban_format: "18n", code: "RS", iban_fields: "RSkk bbbc cccc cccc cccc xx", comment: "b = National bank code c = Account number x = Account check digits", standard_treatment: true},
		"SK": iban_country{country: "Slovakia", chars: 24, bban_format: "20n", code: "SK", iban_fields: "SKkk bbbb ssss sscc cccc cccc", comment: "b = National bank code s = Account number prefix c = Account number", standard_treatment: true},
		"SI": iban_country{country: "Slovenia", chars: 19, bban_format: "15n", code: "SI", iban_fields: "SIkk bbss sccc cccc cxx", comment: "k = IBAN check digits (always = '56') b = National bank code s = Branch code c = Account number x = National check digits", standard_treatment: true},
		"ES": iban_country{country: "Spain", chars: 24, bban_format: "20n", code: "ES", iban_fields: "ESkk bbbb gggg xxcc cccc cccc", comment: "b = National bank code g = Branch code x = Check digits c = Account number", standard_treatment: true},
		"SE": iban_country{country: "Sweden", chars: 24, bban_format: "20n", code: "SE", iban_fields: "SEkk bbbc cccc cccc cccc cccc", comment: "b = National bank code c = Account number ", standard_treatment: true},
		"CH": iban_country{country: "Switzerland", chars: 21, bban_format: "5n,12c", code: "CH", iban_fields: "CHkk bbbb bccc cccc cccc c", comment: "b = National bank code c = Account number", standard_treatment: true},
		"TN": iban_country{country: "Tunisia", chars: 24, bban_format: "20n", code: "TN", iban_fields: "TNkk bbss sccc cccc cccc cccc", comment: "k = IBAN check digits (always 59) b = National bank code s = Branch code c = Account number", standard_treatment: true},
		"TR": iban_country{country: "Turkey", chars: 26, bban_format: "5n,17c", code: "TR", iban_fields: "TRkk bbbb bxcc cccc cccc cccc cc", comment: "b = National bank code x = Reserved for future use (currently '0') c = Account number", standard_treatment: true},
		"GB": iban_country{country: "United Kingdom", chars: 22, bban_format: "4a,14n", code: "GB", iban_fields: "GBkk bbbb ssss sscc cccc cc", comment: "b = BIC bank code s = Bank and branch code (sort code) c = Account number", standard_treatment: true},
		"AE": iban_country{country: "United Arab Emirates", chars: 23, bban_format: "3n,16n", code: "AE", iban_fields: "AEkk bbbc cccc cccc cccc ccc", comment: "b = National bank code c = Account number ", standard_treatment: true},
		"VG": iban_country{country: "Virgin Islands, British", chars: 24, bban_format: "4c,16n", code: "VG", iban_fields: "VGkk bbbb cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standard_treatment: true},
	}
)

//****************** IsCorrectIban **********************************************//
func IsCorrectIban(iban string, debug bool) (bool, error) {
	var iban_config iban_country
	iban_valid := false

	if len(iban) > 15 { // Minimum length for IBAN

		// Clean up string
		iban = strings.ToUpper(strings.Replace(iban, " ", "", -1))

		// Split string up
		passed_chars = len(iban)
		passed_code, passed_checksum, passeds_bban = split_iban_up(iban)

		//fmt.Println(country_list)

		iban_config = country_list[passed_code]

		if iban_config.chars > 0 {

			if debug {
				fmt.Println("--------------------------")
				iban_config.WriteCountryInfo()
				fmt.Println("--Passed Values--")
				fmt.Println("Country code : " + passed_code)
				fmt.Println("Length       : " + strconv.Itoa(passed_chars))
				fmt.Println("BBAN         : " + passeds_bban)
				fmt.Println("Checksum     : " + passed_checksum)
				fmt.Println("--------------------------")
			}
			if iban_config.chars == passed_chars {
				converted_iban := rearrange(passed_code, passed_checksum, passeds_bban)
				converted_iban = convert_char_to_number(converted_iban)

				if calculate_modulo(converted_iban) == 1 {
					iban_valid = true
				}
			} else {
				return false, fmt.Errorf("IBAN: lenght (%s) does not match configuration length (%s)!", strconv.Itoa(passed_chars), strconv.Itoa(iban_config.chars))
			}

		} else {
			return false, fmt.Errorf("IBAN: country <%s> is not in the list!", passed_code)
		}

	} else {
		return false, fmt.Errorf("IBAN: incorect IBAN string passed <%s>", iban)
	}
	return iban_valid, nil
}

func split_iban_up(iban string) (country_code string, checksum string, bban string) {
	country_code = iban[0:2]
	checksum = iban[2:4]
	bban = iban[4:len(iban)]
	return country_code, checksum, bban
}

//****************** GetIbanChecksum **********************************************
func GetIbanChecksum(iban string) (int, error) {
	iban_checksum := 0

	if len(iban) > 15 { // Minimum length for IBAN

		iban = strings.ToUpper(strings.Replace(iban, " ", "", -1))

		// Split string up
		passed_chars = len(iban)
		passed_code, passed_checksum, passeds_bban = split_iban_up(iban)
		passed_checksum = "00"
		converted_iban := rearrange(passed_code, passed_checksum, passeds_bban)
		converted_iban = convert_char_to_number(converted_iban)

		iban_checksum = 98 - calculate_modulo(converted_iban)

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

func rearrange(country_code string, checksum string, bban string) string {
	return bban + country_code + checksum
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
