package vies

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, code, result, expected string) {
	if result != expected {
		t.Errorf("Wrong %s, expected '%s', got '%s'", code, expected, result)
	}
}

func TestWriteFunction(t *testing.T) {
	vatDetail := VatDetails{
		CountryCode: "TEST",
		VatNumber:   "1234567890",
		Address:     "Test Address",
		Error:       "",
		ErrorCode:   "",
		Name:        "Test name",
		Valid:       true,
	}
	writer, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}

	vatDetail.Print(writer)
	writer.Seek(0, os.SEEK_SET)

	bytesResult, _ := ioutil.ReadAll(writer)
	result := strings.Split(string(bytesResult), "\n")

	assertEqual(t, "Error", result[0], "Error: ")
	assertEqual(t, "ErrorCode", result[1], "ErrorCode: ")
	assertEqual(t, "CountryCode", result[2], "CountryCode: TEST")
	assertEqual(t, "VatNumber", result[3], "VatNumber: 1234567890")
	assertEqual(t, "RequestDate", result[4], "RequestDate: ")
	assertEqual(t, "Valid", result[5], "Valid: true")
	assertEqual(t, "Name", result[6], "Name: Test name")
	assertEqual(t, "Address", result[7], "Address: Test Address")
}
