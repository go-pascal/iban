# iban
GoLang IBAN Validation module

The separate BBAN ckeck is not yet implemented.
This is needed for the countries where the IBAN ckecksum is a constant.
Currently the validation will work but the constant is not checked and neiter the BBAN.
Concerned countries:
Macedonia, Bosnia and Herzegovina, East Timor, Mauritania, Montenegro, Portugal, Slovenia, Tunisia

The IBAN may be formated with spaces. Letter cases are ignored.

Here a small sample to test it.


```go
package main

import (
	"fmt"
	"github.com/go-pascal/iban"
)

func main() {

	var OK, well_formated, err = iban.IsCorrectIban("GB82 WEST 1234 5698 7654 32 ", true) // passed: IBAN string , debug true/false
	if err != nil {
		fmt.Println(err.Error())
	} else {

		if OK {
			fmt.Printf("IBAN %s is valid \n\r", well_formated)
		} else {
			fmt.Println("IBAN is not valid")
		}
	}

	var ck, err2 = iban.GetIbanChecksum("GB00 WEST 1234 5698 7654 32 ")
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {

		fmt.Println("IBAN checksum :", ck)
	}
}
```

```go
package main

import "github.com/RitchieFlick/iban"
import "log"

func main() {
	iban, err := iban.NewIBAN("GB82WEST12345698765432")
	if err != nil {
		log.Print(err)
	}
	log.Print(iban.Number)
}

```
