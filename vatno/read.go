package vatno

import (
	"fmt"
	"io"
)

// VATNo describe a basic information about company, such as country code and vat number
type VATNo struct {
	// The country where the company has headquater
	CountryCode string
	// The VAT number of company
	VATNumber string
}

// Function creates VATNo object from reader.
// Reader should prepare a data in two lines, in first one should be the country code,
// in second one the vat number.
// Function returns the VATNo object and any write error encounter.
func Read(reader io.Reader) (VATNo, error) {
	var data VATNo

	_, err := fmt.Fscanf(reader, "%s", &data.CountryCode)
	if err != nil {
		return data, err
	}

	_, err = fmt.Fscanf(reader, "%s", &data.VATNumber)
	if err != nil {
		return data, err
	}

	return data, nil
}
