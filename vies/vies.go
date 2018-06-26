package vies

import (
	"bytes"
	"io"
	"net/http"

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
	var body *bytes.Buffer
	return body
}

func sendViesPostRequest(body io.Reader) (*http.Response, error) {
	res, err := http.Post(viesServiceUrl, "application/xml", body)
	return res, err
}
