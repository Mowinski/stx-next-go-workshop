package vies

import (
	"bytes"
	"testing"

	"github.com/Mowinski/stx-next-go-workshop/vatno"
)

func TestPrepareViesRequestBody(t *testing.T) {
	testData := vatno.VATNo{
		CountryCode: "TEST",
		VATNumber:   "1234567890",
	}
	expectedBody := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><SOAP-ENV:Envelope \txmlns:ns0=\"urn:ec.europa.eu:taxud:vies:services:checkVat:types\"\txmlns:ns1=\"http://schemas.xmlsoap.org/soap/envelope/\" \txmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" \txmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\">\t<SOAP-ENV:Header/>\t<ns1:Body>\t\t<ns0:checkVat>\t\t\t<ns0:countryCode>TEST</ns0:countryCode>\t\t\t<ns0:vatNumber>1234567890</ns0:vatNumber>\t\t</ns0:checkVat>\t</ns1:Body></SOAP-ENV:Envelope>"

	body := prepareViesRequestBody(testData)
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	assertEqual(t, "Body", buf.String(), expectedBody)
}
