package vies

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Mowinski/stx-next-go-workshop/vatno"
)

const viesReqPayloadTmpl = `
<?xml version="1.0" encoding="UTF-8"?><SOAP-ENV:Envelope 
	xmlns:ns0="urn:ec.europa.eu:taxud:vies:services:checkVat:types"	xmlns:ns1="http://schemas.xmlsoap.org/soap/envelope/" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
	<SOAP-ENV:Header/>
	<ns1:Body>
		<ns0:checkVat>
			<ns0:countryCode>%s</ns0:countryCode>
			<ns0:vatNumber>%s</ns0:vatNumber>
		</ns0:checkVat>
	</ns1:Body>
</SOAP-ENV:Envelope>`

const viesServiceUrl = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"

func prepareViesRequestBody(vatDetail vatno.VATNo) io.Reader {
	req := fmt.Sprintf(
		viesReqPayloadTmpl,
		vatDetail.CountryCode,
		vatDetail.VATNumber,
	)
	body := bytes.NewBufferString(
		strings.Replace(req, "\n", "", -1),
	)
	return body
}

func sendViesPostRequest(body io.Reader) (*http.Response, error) {
	res, err := http.Post(viesServiceUrl, "application/xml", body)
	return res, err
}

func decodeResponse(r io.Reader) (VatDetails, error) {
	var vatDetails VatDetails

	enc := xml.NewDecoder(r)
	err := enc.Decode(&vatDetails)
	return vatDetails, err
}

// Query function retrieves company details from vias service.
// Input data is a struct with country code and vat number.
// If the company exists and is registered as VAT client, the function returns
// a struct with detailed information, e.g: company name and address.
// Otherwise the function returns the struct with field Valid as False.
// It returns pointer to the VatDetails struct and any error encounter.
func Query(vatNo vatno.VATNo) (*VatDetails, error) {
	body := prepareViesRequestBody(vatNo)
	response, err := sendViesPostRequest(body)
	if err != nil {
		return nil, err
	}
	vatDetails, err := decodeResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return &vatDetails, nil
}
