package vies

import (
	"strings"
	"testing"
)

func assertEqual(t *testing.T, code, result, expected string) {
	if result != expected {
		t.Errorf("Wrong %s, expected '%s', got '%s'", code, expected, result)
	}
}

func assertBoolEqual(t *testing.T, code string, result, expected bool) {
	if result != expected {
		t.Errorf("Wrong %s, expected '%t', got '%t'", code, expected, result)
	}
}

func TestPrepareViesRequestBody(t *testing.T) {
	testData := InputData{
		CountryCode: "TEST",
		VATNumber:   "1234567890",
	}
	expectedBody := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><SOAP-ENV:Envelope \txmlns:ns0=\"urn:ec.europa.eu:taxud:vies:services:checkVat:types\"\txmlns:ns1=\"http://schemas.xmlsoap.org/soap/envelope/\" \txmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" \txmlns:SOAP-ENV=\"http://schemas.xmlsoap.org/soap/envelope/\">\t<SOAP-ENV:Header/>\t<ns1:Body>\t\t<ns0:checkVat>\t\t\t<ns0:countryCode>TEST</ns0:countryCode>\t\t\t<ns0:vatNumber>1234567890</ns0:vatNumber>\t\t</ns0:checkVat>\t</ns1:Body></SOAP-ENV:Envelope>"

	body := prepareViesRequestBody(testData)

	assertEqual(t, "Body", body.String(), expectedBody)
}

func TestEncodeResponseFromVies(t *testing.T) {
	body := `
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body>
		<checkVatResponse
			xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
			<countryCode>TEST</countryCode>
			<vatNumber>1234567890</vatNumber>
			<requestDate>2018-06-16+02:00</requestDate>
			<valid>true</valid>
			<name>Test Name</name>
			<address>Test Address</address>
		</checkVatResponse>
	</soap:Body>
</soap:Envelope>`
	bodyReader := strings.NewReader(body)

	result, err := encodeResponseFromVies(bodyReader)

	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Error", result.Error, "")
	assertEqual(t, "ErrorCode", result.ErrorCode, "")
	assertEqual(t, "CountryCode", result.CountryCode, "TEST")
	assertEqual(t, "VatNumber", result.VatNumber, "1234567890")
	assertEqual(t, "RequestDate", result.RequestDate, "2018-06-16+02:00")
	assertEqual(t, "Name", result.Name, "Test Name")
	assertEqual(t, "Address", result.Address, "Test Address")
	assertBoolEqual(t, "Valid", result.Valid, true)

}

func TestCheck(t *testing.T) {
	testData := InputData{
		CountryCode: "GB",
		VATNumber:   "333289454",
	}

	result, err := Check(testData)

	if err != nil {
		t.Fatal(err)
	}

	assertBoolEqual(t, "Valid", result.Valid, true)
	assertEqual(t, "CountryCode", result.CountryCode, "GB")
	assertEqual(t, "VatNumber", result.VatNumber, "333289454")
	assertEqual(t, "Name", result.Name, "BRITISH BROADCASTING CORPORATION")
	assertEqual(t, "Error", result.Error, "")
	assertEqual(t, "ErrorCode", result.ErrorCode, "")
}

func TestCheckWithWrongVatNumber(t *testing.T) {
	testData := InputData{
		CountryCode: "GB",
		VATNumber:   "012345678",
	}

	result, err := Check(testData)

	if err != nil {
		t.Fatal(err)
	}

	assertBoolEqual(t, "Valid", result.Valid, false)
	assertEqual(t, "CountryCode", result.CountryCode, "GB")
	assertEqual(t, "VatNumber", result.VatNumber, "012345678")
	assertEqual(t, "Name", result.Name, "---")
	assertEqual(t, "Error", result.Error, "")
	assertEqual(t, "ErrorCode", result.ErrorCode, "")
}
