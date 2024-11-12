package iban

var countryList = map[string]ibanCountry{
	"AD": {country: "Andorra", chars: 24, bbanFormat: "8n,12c", code: "AD", ibanFields: "ADkk bbbb ssss cccc cccc cccc", comment: "b = National bank code s = Branch code c = Account number", standardTreatment: true},
	"AE": {country: "United Arab Emirates", chars: 23, bbanFormat: "3n,16n", code: "AE", ibanFields: "AEkk bbbc cccc cccc cccc ccc", comment: "b = National bank code c = Account number ", standardTreatment: true},
	"AL": {country: "Albania", chars: 28, bbanFormat: "8n, 16c", code: "AL", ibanFields: "ALkk bbbs sssx cccc cccc cccc cccc", comment: "b = National bank code s = Branch code x = National check digit c = Account number", standardTreatment: true},
	"AO": {country: "Angola", chars: 25, bbanFormat: "21n", code: "AO", ibanFields: "AOkk bbbb cccc cccc cccc cccx", comment: "b = Bank code; c = Account number; x = Check digit", standardTreatment: true},
	"AT": {country: "Austria", chars: 20, bbanFormat: "16n", code: "AT", ibanFields: "ATkk bbbb bccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"AZ": {country: "Azerbaijan", chars: 28, bbanFormat: "4c,20n", code: "AZ", ibanFields: "AZkk bbbb cccc cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"BA": {country: "Bosnia and Herzegovina", chars: 20, bbanFormat: "16n", code: "BA", ibanFields: "BAkk bbbs sscc cccc ccxx", comment: "k = IBAN check digits (always 39) b = National bank code s = Branch code c = Account number x = National check digits", standardTreatment: true},
	"BE": {country: "Belgium", chars: 16, bbanFormat: "12n", code: "BE", ibanFields: "BEkk bbbc cccc ccxx", comment: "b = National bank code c = Account number x = National check digits", standardTreatment: true},
	"BF": {country: "Burkina Faso", chars: 28, bbanFormat: "24n", code: "BF", ibanFields: "BFkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"BG": {country: "Bulgaria", chars: 22, bbanFormat: "4a,6n,8c", code: "BG", ibanFields: "BGkk bbbb ssss ddcc cccc cc", comment: "b = BIC bank code s = Branch (BAE) number d = Account type c = Account number", standardTreatment: true},
	"BH": {country: "Bahrain", chars: 22, bbanFormat: "4a,14c", code: "BH", ibanFields: "BHkk bbbb cccc cccc cccc cc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"BI": {country: "Burundi", chars: 16, bbanFormat: "12n", code: "BI", ibanFields: "BIkk bbbb cccc cccc", comment: "b = Bank code; c = Account number", standardTreatment: true},
	"BJ": {country: "Benin", chars: 28, bbanFormat: "24n", code: "BJ", ibanFields: "BJkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"BL": {country: "Saint Barthélemy", chars: 27, bbanFormat: "10n,11c,2n", code: "BL", ibanFields: "BLkk bbbb bsss sscc cccc cccc cxx", comment: "b = Bank code; s = Branch code; c = Account number; x = National check digits", standardTreatment: true},
	"BR": {country: "Brazil", chars: 29, bbanFormat: "23n, 1a, 1c", code: "BR", ibanFields: "BRkk bbbb bbbb ssss sccc cccc ccct n", comment: "k = IBAN check digits (Calculated by MOD 97-10) b = National bank code s = Branch code c = Account Number t = Account type (Cheque account, Savings account etc.) n = Owner account number (1, 2 etc.)[31]", standardTreatment: true},
	"BY": {country: "Belarus", chars: 28, bbanFormat: "4c,20n", code: "BY", ibanFields: "BYkk bbbb cccc cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"CF": {country: "Central African Republic", chars: 27, bbanFormat: "23n", code: "CF", ibanFields: "CFkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"CG": {country: "Congo", chars: 27, bbanFormat: "23n", code: "CG", ibanFields: "CGkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"CH": {country: "Switzerland", chars: 21, bbanFormat: "5n,12c", code: "CH", ibanFields: "CHkk bbbb bccc cccc cccc c", comment: "b = National bank code c = Account number", standardTreatment: true},
	"CI": {country: "Ivory Coast", chars: 28, bbanFormat: "24n", code: "CI", ibanFields: "CIkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"CM": {country: "Cameroon", chars: 27, bbanFormat: "23n", code: "CM", ibanFields: "CMkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"CR": {country: "Costa Rica", chars: 22, bbanFormat: "17n", code: "CR", ibanFields: "CRkk bbbc cccc cccc cccc c", comment: "b = bank code c = Account number", standardTreatment: false},
	"CV": {country: "Cape Verde", chars: 25, bbanFormat: "21n", code: "CV", ibanFields: "CVkk bbbb ssss cccc cccc cccx", comment: "b = Bank code; s = Branch code; c = Account number; x = Check digit", standardTreatment: true},
	"CY": {country: "Cyprus", chars: 28, bbanFormat: "8n,16c", code: "CY", ibanFields: "CYkk bbbs ssss cccc cccc cccc cccc", comment: "b = National bank code s = Branch code c = Account number", standardTreatment: true},
	"CZ": {country: "Czech Republic", chars: 24, bbanFormat: "20n", code: "CZ", ibanFields: "CZkk bbbb ssss sscc cccc cccc", comment: "b = National bank code s = Account number prefix c = Account number", standardTreatment: true},
	"DE": {country: "Germany", chars: 22, bbanFormat: "18n", code: "DE", ibanFields: "DEkk bbbb bbbb cccc cccc cc", comment: "b = Bank and branch identifier (de:Bankleitzahl or BLZ) c = Account number", standardTreatment: true},
	"DJ": {country: "Djibouti", chars: 27, bbanFormat: "23n", code: "DJ", ibanFields: "DJkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"DK": {country: "Denmark", chars: 18, bbanFormat: "14n", code: "DK", ibanFields: "DKkk bbbb cccc cccc cc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"DO": {country: "Dominican Republic", chars: 28, bbanFormat: "4a,20n", code: "DO", ibanFields: "DOkk bbbb cccc cccc cccc cccc cccc", comment: "b = Bank identifier c = Account number", standardTreatment: true},
	"DZ": {country: "Algeria", chars: 26, bbanFormat: "20n", code: "DZ", ibanFields: "DZkk bbbb ssss cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"EE": {country: "Estonia", chars: 20, bbanFormat: "16n", code: "EE", ibanFields: "EEkk bbss cccc cccc cccx", comment: "b = National bank code s = Branch code c = Account number x = National check digit", standardTreatment: true},
	"EG": {country: "Egypt", chars: 29, bbanFormat: "25n", code: "EG", ibanFields: "EGkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"ES": {country: "Spain", chars: 24, bbanFormat: "20n", code: "ES", ibanFields: "ESkk bbbb gggg xxcc cccc cccc", comment: "b = National bank code g = Branch code x = Check digits c = Account number", standardTreatment: true},
	"FI": {country: "Finland", chars: 18, bbanFormat: "14n", code: "FI", ibanFields: "FIkk bbbb bbcc cccc cx", comment: "b = Bank and branch code c = Account number x = National check digit", standardTreatment: true},
	"FO": {country: "Faroe Islands", chars: 18, bbanFormat: "14n", code: "FO", ibanFields: "FOkk bbbb cccc cccc cx", comment: "b = National bank code c = Account number x = National check digit", standardTreatment: true},
	"FR": {country: "France", chars: 27, bbanFormat: "10n,11c,2n", code: "FR", ibanFields: "FRkk bbbb bggg ggcc cccc cccc cxx", comment: "b = National bank code g = Branch code (fr:code guichet) c = Account number x = National check digits (fr:clé RIB)", standardTreatment: true},
	"GA": {country: "Gabon", chars: 27, bbanFormat: "23n", code: "GA", ibanFields: "GAkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"GB": {country: "United Kingdom", chars: 22, bbanFormat: "4a,14n", code: "GB", ibanFields: "GBkk bbbb ssss sscc cccc cc", comment: "b = BIC bank code s = Bank and branch code (sort code) c = Account number", standardTreatment: true},
	"GE": {country: "Georgia", chars: 22, bbanFormat: "2c,16n", code: "GE", ibanFields: "GEkk bbcc cccc cccc cccc cc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"GG": {country: "Guernsey", chars: 22, bbanFormat: "4a,14n", code: "GG", ibanFields: "GGkk bbbb ssss sscc cccc cc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"GI": {country: "Gibraltar", chars: 23, bbanFormat: "4a,15c", code: "GI", ibanFields: "GIkk bbbb cccc cccc cccc ccc", comment: "b = BIC bank code c = Account number", standardTreatment: true},
	"GL": {country: "Greenland", chars: 18, bbanFormat: "14n", code: "GL", ibanFields: "GLkk bbbb cccc cccc cc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"GQ": {country: "Equatorial Guinea", chars: 27, bbanFormat: "23n", code: "GQ", ibanFields: "GQkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"GR": {country: "Greece", chars: 27, bbanFormat: "7n,16c", code: "GR", ibanFields: "GRkk bbbs sssc cccc cccc cccc ccc", comment: "b = National bank code s = Branch code c = Account number", standardTreatment: true},
	"GT": {country: "Guatemala", chars: 28, bbanFormat: "4c,20c", code: "GT", ibanFields: "GTkk bbbb mmtt cccc cccc cccc cccc", comment: "b = National bank code c = Account number m = Currency t = Account type ", standardTreatment: true},
	"GW": {country: "Guinea Bissau", chars: 25, bbanFormat: "21n", code: "GW", ibanFields: "GWkk bbbb ssss cccc cccc cccx", comment: "b = Bank code; s = Branch code; c = Account number; x = Check digit", standardTreatment: true},
	"HN": {country: "Honduras", chars: 28, bbanFormat: "24n", code: "HN", ibanFields: "HNkk pppp cccc cccc cccc cccc cccc", comment: "p = Bank identifier code; c = Account number", standardTreatment: true},
	"HR": {country: "Croatia", chars: 21, bbanFormat: "17n", code: "HR", ibanFields: "HRkk bbbb bbbc cccc cccc c", comment: "b = Bank code c = Account number", standardTreatment: true},
	"HU": {country: "Hungary", chars: 28, bbanFormat: "24n", code: "HU", ibanFields: "HUkk bbbs sssk cccc cccc cccc cccx", comment: "b = National bank code s = Branch code c = Account number x = National check digit", standardTreatment: true},
	"IE": {country: "Ireland", chars: 22, bbanFormat: "4c,14n", code: "IE", ibanFields: "IEkk aaaa bbbb bbcc cccc cc", comment: "a = BIC bank code b = Bank/branch code (sort code) c = Account number", standardTreatment: true},
	"IK": {country: "Israel", chars: 23, bbanFormat: "19n", code: "IK", ibanFields: "ILkk bbbn nncc cccc cccc ccc", comment: "b = National bank code n = Branch number c = Account number 13 digits (padded with zeros)", standardTreatment: true},
	"IL": {country: "Israel", chars: 23, bbanFormat: "4c,15n", code: "IL", ibanFields: "ILkk bbbb cccc cccc cccc ccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"IM": {country: "Isle of Man", chars: 22, bbanFormat: "4a,14n", code: "IM", ibanFields: "IMkk bbbb ssss sscc cccc cc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"IQ": {country: "Iraq", chars: 23, bbanFormat: "4c,15n", code: "IQ", ibanFields: "IQkk bbbb cccc cccc cccc ccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"IR": {country: "Iran", chars: 26, bbanFormat: "22n", code: "IR", ibanFields: "IRkk pppp cccc cccc cccc cccc cc", comment: "p = Bank code; c = Account number", standardTreatment: true},
	"IS": {country: "Iceland", chars: 26, bbanFormat: "22n", code: "IS", ibanFields: "ISkk bbbb sscc cccc iiii iiii ii", comment: "b = National bank code s = Branch code c = Account number i = holder's kennitala (national identification number).", standardTreatment: true},
	"IT": {country: "Italy", chars: 27, bbanFormat: "1a,10n,12c", code: "IT", ibanFields: "ITkk xaaa aabb bbbc cccc cccc ccc", comment: "x = Check char (CIN) a = National bank code (it:Associazione bancaria italiana or Codice ABI ) b = Branch code (it:Coordinate bancarie or CAB – Codice d'Avviamento Bancario) c = Account number", standardTreatment: true},
	"JE": {country: "Jersey", chars: 22, bbanFormat: "4a,14n", code: "JE", ibanFields: "JEkk bbbb ssss sscc cccc cc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"JO": {country: "Jordan", chars: 30, bbanFormat: "4a, 22n", code: "JO", ibanFields: "JOkk bbbb nnnn cccc cccc cccc cccc cc", comment: "b = National bank code n = Branch code c = Account number ", standardTreatment: true},
	"KM": {country: "Comoros", chars: 27, bbanFormat: "23n", code: "KM", ibanFields: "KMkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"KW": {country: "Kuwait", chars: 30, bbanFormat: "4a, 22c", code: "KW", ibanFields: "KWkk bbbb cccc cccc cccc cccc cccc cc", comment: "b = National bank code c = Account number.", standardTreatment: true},
	"KZ": {country: "Kazakhstan", chars: 20, bbanFormat: "3n,13c", code: "KZ", ibanFields: "KZkk bbbc cccc cccc cccc", comment: "b = National bank code c = Account number ", standardTreatment: true},
	"LB": {country: "Lebanon", chars: 28, bbanFormat: "4n,20c", code: "LB", ibanFields: "LBkk bbbb cccc cccc cccc cccc cccc", comment: "b = Bank code; c = Account number", standardTreatment: true},
	"LC": {country: "Saint Lucia", chars: 32, bbanFormat: "4c,24n", code: "LC", ibanFields: "LCkk bbbb cccc cccc cccc ccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"LI": {country: "Liechtenstein", chars: 21, bbanFormat: "5n,12c", code: "LI", ibanFields: "LIkk bbbb bccc cccc cccc c", comment: "b = National bank code c = Account number", standardTreatment: true},
	"LT": {country: "Lithuania", chars: 20, bbanFormat: "16n", code: "LT", ibanFields: "LTkk bbbb bccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"LU": {country: "Luxembourg", chars: 20, bbanFormat: "3n,13c", code: "LU", ibanFields: "LUkk bbbc cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"LV": {country: "Latvia", chars: 21, bbanFormat: "4a,13c", code: "LV", ibanFields: "LVkk bbbb cccc cccc cccc c", comment: "b = BIC Bank code c = Account number", standardTreatment: true},
	"MA": {country: "Morocco", chars: 28, bbanFormat: "24n", code: "MA", ibanFields: "MAkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"MC": {country: "Monaco", chars: 27, bbanFormat: "10n,11c,2n", code: "MC", ibanFields: "MCkk bbbb bsss sscc cccc cccc cxx", comment: "b = National bank code s = Branch code (fr:code guichet) c = Account number x = National check digits (fr:clé RIB). ", standardTreatment: true},
	"MD": {country: "Moldova", chars: 24, bbanFormat: "2c,18c", code: "MD", ibanFields: "MDkk bbcc cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"ME": {country: "Montenegro", chars: 22, bbanFormat: "18n", code: "ME", ibanFields: "MEkk bbbc cccc cccc cccc xx", comment: "k = IBAN check digits (always = '25') b = Bank code c = Account number x = National check digits", standardTreatment: true},
	"MF": {country: "Saint Martin", chars: 27, bbanFormat: "10n,11c,2n", code: "MF", ibanFields: "MFkk bbbb bsss sscc cccc cccc cxx", comment: "b = Bank code; s =  Branch code; c = Account number; x = National check digits", standardTreatment: true},
	"MG": {country: "Madagascar", chars: 27, bbanFormat: "23n", code: "MG", ibanFields: "MGkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"MK": {country: "Macedonia", chars: 19, bbanFormat: "3n,10c,2n", code: "MK", ibanFields: "MKkk bbbc cccc cccc cxx", comment: "k = IBAN check digits (always = '07') b = National bank code c = Account number x = National check digits", standardTreatment: true},
	"ML": {country: "Mali", chars: 28, bbanFormat: "24n", code: "ML", ibanFields: "MLkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"MR": {country: "Mauritania", chars: 27, bbanFormat: "23n", code: "MR", ibanFields: "MRkk bbbb bsss sscc cccc cccc cxx", comment: "k = IBAN check digits (always 13) b = National bank code s = Branch code (fr:code guichet) c = Account number x = National check digits (fr:clé RIB)", standardTreatment: true},
	"MT": {country: "Malta", chars: 31, bbanFormat: "4a,5n,18c", code: "MT", ibanFields: "MTkk bbbb ssss sccc cccc cccc cccc ccc", comment: "b = BIC bank code s = Branch code c = Account number", standardTreatment: true},
	"MU": {country: "Mauritius", chars: 30, bbanFormat: "4a,19n,3a", code: "MU", ibanFields: "MUkk bbbb bbss cccc cccc cccc 000d dd", comment: "b = National bank code s = Branch identifier c = Account number 0 = Zeroes d = Currency Symbol ", standardTreatment: true},
	"MZ": {country: "Mozambique", chars: 25, bbanFormat: "21n", code: "MZ", ibanFields: "MZkk bbbb ssss cccc cccc cccx x", comment: "b = Bank code; s = Branch code; c = Account number; x = Check digit", standardTreatment: true},
	"NE": {country: "Niger", chars: 28, bbanFormat: "24n", code: "NE", ibanFields: "NEkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"NI": {country: "Nicaragua", chars: 32, bbanFormat: "28n", code: "NI", ibanFields: "NIkk bbbb ssss cccc cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"NL": {country: "Netherlands", chars: 18, bbanFormat: "4a,10n", code: "NL", ibanFields: "NLkk bbbb cccc cccc cc", comment: "b = BIC Bank code c = Account number", standardTreatment: true},
	"NO": {country: "Norway", chars: 15, bbanFormat: "11n", code: "NO", ibanFields: "NOkk bbbb cccc ccx", comment: "b = National bank code c = Account number x = Modulo-11 national check digit", standardTreatment: true},
	"PK": {country: "Pakistan", chars: 24, bbanFormat: "4c,16n", code: "PK", ibanFields: "PKkk bbbb cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"PL": {country: "Poland", chars: 28, bbanFormat: "24n", code: "PL", ibanFields: "PLkk bbbs sssx cccc cccc cccc cccc", comment: "b = National bank code s = Branch code x = National check digit c = Account number, ", standardTreatment: true},
	"PS": {country: "Palestinian territories", chars: 29, bbanFormat: "4c,21n", code: "PS", ibanFields: "PSkk bbbb xxxx xxxx xccc cccc cccc c", comment: "b = National bank code c = Account number x = Not specified", standardTreatment: true},
	"PT": {country: "Portugal", chars: 25, bbanFormat: "21n", code: "PT", ibanFields: "PTkk bbbb ssss cccc cccc cccx x", comment: "k = IBAN check digits (always = '50') b = National bank code s = Branch code C = Account number x = National check digit", standardTreatment: true},
	"QA": {country: "Qatar", chars: 29, bbanFormat: "4a, 21c", code: "QA", ibanFields: "QAkk bbbb cccc cccc cccc cccc cccc c", comment: "b = National bank code c = Account number[34]", standardTreatment: true},
	"RE": {country: "Réunion", chars: 27, bbanFormat: "10n,11c,2n", code: "RE", ibanFields: "REkk bbbb bsss sscc cccc cccc cxx", comment: "b = Bank code; s = Branch code; c = Account number; x = National check digits", standardTreatment: true},
	"RO": {country: "Romania", chars: 24, bbanFormat: "4a,16c", code: "RO", ibanFields: "ROkk bbbb cccc cccc cccc cccc", comment: "b = BIC Bank code c = Branch code and account number (bank-specific format) ", standardTreatment: true},
	"RS": {country: "Serbia", chars: 22, bbanFormat: "18n", code: "RS", ibanFields: "RSkk bbbc cccc cccc cccc xx", comment: "b = National bank code c = Account number x = Account check digits", standardTreatment: true},
	"SA": {country: "Saudi Arabia", chars: 24, bbanFormat: "2n,18c", code: "SA", ibanFields: "SAkk bbcc cccc cccc cccc cccc", comment: "b = National bank code c = Account number preceded by zeros, if required", standardTreatment: true},
	"SC": {country: "Seychelles", chars: 31, bbanFormat: "4c,23n", code: "SC", ibanFields: "SCkk bbbb cccc cccc cccc cccc cccc mmm", comment: "b = National bank code c = Account number m = Currency", standardTreatment: true},
	"SE": {country: "Sweden", chars: 24, bbanFormat: "20n", code: "SE", ibanFields: "SEkk bbbc cccc cccc cccc cccc", comment: "b = National bank code c = Account number ", standardTreatment: true},
	"SI": {country: "Slovenia", chars: 19, bbanFormat: "15n", code: "SI", ibanFields: "SIkk bbss sccc cccc cxx", comment: "k = IBAN check digits (always = '56') b = National bank code s = Branch code c = Account number x = National check digits", standardTreatment: true},
	"SK": {country: "Slovakia", chars: 24, bbanFormat: "20n", code: "SK", ibanFields: "SKkk bbbb ssss sscc cccc cccc", comment: "b = National bank code s = Account number prefix c = Account number", standardTreatment: true},
	"SM": {country: "San Marino", chars: 27, bbanFormat: "1a,10n,12c", code: "SM", ibanFields: "SMkk xaaa aabb bbbc cccc cccc ccc", comment: "x = Check char (it:CIN) a = National bank code (it:Associazione bancaria italiana or Codice ABI) b = Branch code (it:Coordinate bancarie or CAB – Codice d'Avviamento Bancario) c = Account number", standardTreatment: true},
	"SN": {country: "Senegal", chars: 28, bbanFormat: "24n", code: "SN", ibanFields: "SNkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"ST": {country: "Sao Tome and Principe", chars: 25, bbanFormat: "4c,17n", code: "ST", ibanFields: "STkk bbbb cccc cccc cccc cccc c", comment: "b = National bank code c = Account number", standardTreatment: true},
	"SV": {country: "El Salvador", chars: 28, bbanFormat: "4c,20n", code: "SV", ibanFields: "SVkk bbbb cccc cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"TD": {country: "Chad", chars: 27, bbanFormat: "23n", code: "TD", ibanFields: "TDkk bbbb ssss cccc cccc cccc ccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"TG": {country: "Togo", chars: 28, bbanFormat: "24n", code: "TG", ibanFields: "TGkk bbbb ssss cccc cccc cccc cccc", comment: "b = Bank code; s = Branch code; c = Account number", standardTreatment: true},
	"TL": {country: "East Timor", chars: 23, bbanFormat: "19n", code: "TL", ibanFields: "TLkk bbbc cccc cccc cccc cxx", comment: "k = IBAN check digits (always = '38') b = Bank identifier c = Account number x = National check digit", standardTreatment: true},
	"TN": {country: "Tunisia", chars: 24, bbanFormat: "20n", code: "TN", ibanFields: "TNkk bbss sccc cccc cccc cccc", comment: "k = IBAN check digits (always 59) b = National bank code s = Branch code c = Account number", standardTreatment: true},
	"TR": {country: "Turkey", chars: 26, bbanFormat: "5n,17c", code: "TR", ibanFields: "TRkk bbbb bxcc cccc cccc cccc cc", comment: "b = National bank code x = Reserved for future use (currently '0') c = Account number", standardTreatment: true},
	"UA": {country: "Ukraine", chars: 29, bbanFormat: "4c,21n", code: "UA", ibanFields: "UAkk bbbb cccc cccc cccc cccc cccc c", comment: "b = National bank code c = Account number", standardTreatment: true},
	"VA": {country: "Vatican", chars: 22, bbanFormat: "3n,15n", code: "VA", ibanFields: "VAkk bbb cccc cccc cccc ccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"VG": {country: "Virgin Islands, British", chars: 24, bbanFormat: "4c,16n", code: "VG", ibanFields: "VGkk bbbb cccc cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"XK": {country: "Kosovo", chars: 20, bbanFormat: "4n,10n,2n", code: "XK", ibanFields: "XKkk bbbb cccc cccc cccc", comment: "b = National bank code c = Account number", standardTreatment: true},
	"YT": {country: "Mayotte", chars: 27, bbanFormat: "10n,11c,2n", code: "YT", ibanFields: "YTkk bbbb bsss sscc cccc cccc cxx", comment: "b = Bank code; s = Branch code; c = Account number; x = National check digits", standardTreatment: true},
}
