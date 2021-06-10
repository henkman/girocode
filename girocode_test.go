package girocode_test

import (
	"os"
	"testing"

	currency "github.com/bojanz/currency"

	"github.com/henkman/girocode"
)

func Test(t *testing.T) {
	amount, err := currency.NewAmount("2.50", "EUR")
	if err != nil {
		t.Fatal(err)
	}
	transfer := girocode.Transfer{
		Beneficiary: "some person",
		IBAN:        "GB29NWBK60161331926819",
		Amount:      amount,
		Reference:   "for good stuff",
	}
	f, err := os.OpenFile("qr.png", os.O_CREATE|os.O_WRONLY, 0750)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := girocode.Generate(f, transfer); err != nil {
		t.Fatal(err)
	}
}
