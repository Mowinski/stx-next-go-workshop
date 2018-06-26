package vies

// VatDetails describe response from vias service.
type VatDetails struct {
	// the country code of company
	CountryCode string
	// the vat number of company
	VatNumber string
	// the date of request
	RequestDate string
	// describe if company with selected country code and vat number exists
	Valid bool
	// the full company name
	Name string
	// the full company address
	Address string
	// if unexpected error occures, this field describe error message
	Error string
	// if unexpected error occures, this field describe error code
	ErrorCode string
}
