package vies

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type InputData struct {
	CountryCode string
	VATNumber   string
}

type VatDetails struct {
	XMLName     xml.Name `xml:"Envelope"`
	CountryCode string   `xml:"Body>checkVatResponse>countryCode"`
	VatNumber   string   `xml:"Body>checkVatResponse>vatNumber"`
	RequestDate string   `xml:"Body>checkVatResponse>requestDate"`
	Valid       bool     `xml:"Body>checkVatResponse>valid"`
	Name        string   `xml:"Body>checkVatResponse>name"`
	Address     string   `xml:"Body>checkVatResponse>address"`
	Error       string   `xml:"Body>Fault>faultstring"`
	ErrorCode   string   `xml:"Body>Fault>faultcode"`
}

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

func prepareViesRequestBody(vatDetail InputData) *bytes.Buffer {
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

func sendViesPostRequest(body *bytes.Buffer) (*http.Response, error) {
	res, err := http.Post(viesServiceUrl, "application/xml", body)
	return res, err
}

func encodeResponseFromVies(body io.Reader) (vatDetails VatDetails, err error) {
	enc := xml.NewDecoder(body)
	err = enc.Decode(&vatDetails)
	return vatDetails, err
}

func Check(vatDetail InputData) (*VatDetails, error) {
	body := prepareViesRequestBody(vatDetail)
	response, err := sendViesPostRequest(body)
	if err != nil {
		return nil, err
	}
	vatDetails, err := encodeResponseFromVies(response.Body)
	if err != nil {
		return nil, err
	}
	return &vatDetails, nil
}
