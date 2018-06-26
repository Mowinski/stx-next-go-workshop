package vies

import (
	"fmt"
	"io"
)

// Print display the vd object into the writeable w.
// A format of output is:
// 	Error: <ErrorMessage>
// 	ErrorCode: <ErrorCode>
// 	CountryCode: <CountryCode>
// 	VatNumber: <VatNumber>
// 	RequestDate: <RequestDate>
// 	Valid: <ValidFlag>
// 	Name: <CompanyName>
// 	Address: <CompanyAddress>
// It returns any write error encounter
func (vd *VatDetails) Print(w io.Writer) error {
	if _, err := fmt.Fprintln(w, "Error:", vd.Error); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "ErrorCode:", vd.ErrorCode); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "CountryCode:", vd.CountryCode); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "VatNumber:", vd.VatNumber); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "RequestDate:", vd.RequestDate); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "Valid:", vd.Valid); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "Name:", vd.Name); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "Address:", vd.Address); err != nil {
		return err
	}

	return nil
}
