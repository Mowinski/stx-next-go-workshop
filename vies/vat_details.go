package vies

// VatDetails describe response from vias service.
type VatDetails struct {
	// the country code of company
	CountryCode string `xml:"Body>checkVatResponse>countryCode"`
	// the vat number of company
	VatNumber string `xml:"Body>checkVatResponse>vatNumber"`
	// the date of request
	RequestDate string `xml:"Body>checkVatResponse>requestDate"`
	// describe if company with selected country code and vat number exists
	Valid bool `xml:"Body>checkVatResponse>valid"`
	// the full company name
	Name string `xml:"Body>checkVatResponse>name"`
	// the full company address
	Address string `xml:"Body>checkVatResponse>address"`
	// if unexpected error occures, this field describe error message
	Error string `xml:"Body>Fault>faultstring"`
	// if unexpected error occures, this field describe error code
	ErrorCode string `xml:"Body>Fault>faultcode"`
}
