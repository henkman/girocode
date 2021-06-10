package girocode

import (
	"io"
	"strings"

	currency "github.com/bojanz/currency"
	qrcode "github.com/skip2/go-qrcode"
)

type Transfer struct {
	BIC         string
	Beneficiary string
	IBAN        string
	Amount      currency.Amount
	Reference   string
}

func Generate(out io.Writer, t Transfer) error {
	var sb strings.Builder
	sb.WriteString("BCD\n002\n1\nSCT\n")
	if t.BIC != "" {
		sb.WriteString(t.BIC)
	}
	sb.WriteByte('\n')
	sb.WriteString(t.Beneficiary)
	sb.WriteByte('\n')
	sb.WriteString(t.IBAN)
	sb.WriteByte('\n')
	sb.WriteString(t.Amount.CurrencyCode())
	sb.WriteString(t.Amount.Number())
	sb.WriteByte('\n')
	sb.WriteByte('\n')
	sb.WriteString(t.Reference)
	sb.WriteByte('\n')
	qr, err := qrcode.New(sb.String(), qrcode.Medium)
	if err != nil {
		return err
	}
	return qr.Write(256, out)
}
