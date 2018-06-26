package vies

import (
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
	return nil
}
