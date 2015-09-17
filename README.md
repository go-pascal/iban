# iban
GoLang IBAN Validation module

The BBAN ckeck is not yet implemented.
This is needed for teh countries where the IBAN ckecksum is a constant.
Currently the validation will return false.
Concerned countries:
Macedonia, Bosnia and Herzegovina, 

Here a small sample to test it






package main

import (
	"fmt"
	"github.com/go-pascal/iban"
)

func main() {

	var OK, err = iban.IsCorrectIban("GB82 WEST 1234 5698 7654 32 ")
	if err != nil {
		fmt.Println(err.Error())
	} else {

		if OK {
			fmt.Println("IBAN is valid")
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
