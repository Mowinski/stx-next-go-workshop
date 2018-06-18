package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Mowinski/stx-next-go-workshop/vies"
)

func prepareReader(countryCode, number string) (*os.File, error) {
	reader, err := ioutil.TempFile("", "")
	if err != nil {
		return reader, err
	}

	_, err = io.WriteString(reader, countryCode+"\n"+number+"\n")
	if err != nil {
		return reader, err
	}

	_, err = reader.Seek(0, os.SEEK_SET)
	if err != nil {
		return reader, err
	}

	return reader, nil
}

func assertEqual(t *testing.T, code, result, expected string) {
	if result != expected {
		t.Errorf("Wrong %s, expected '%s', got '%s'", code, expected, result)
	}
}

func TestReadingFromFunction(t *testing.T) {
	reader, err := prepareReader("TEST", "1234567890")
	defer reader.Close()

	if err != nil {
		t.Fatal(err)
	}

	result, err := readFrom(reader)

	if err != nil {
		t.Errorf("Error expect %s", err)
	}

	assertEqual(t, "CountryCode", result.CountryCode, "TEST")
	assertEqual(t, "VATNumber", result.VATNumber, "1234567890")
}

func TestWritingToFunction(t *testing.T) {
	vatDetail := vies.VatDetails{
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

	writeTo(writer, &vatDetail)
	writer.Seek(0, os.SEEK_SET)

	bytesResult, _ := ioutil.ReadAll(writer)
	result := strings.Split(string(bytesResult), "\n")

	assertEqual(t, "Error", result[0], "Error: ")
	assertEqual(t, "ErrorCode", result[1], "ErrorCode: ")
	assertEqual(t, "XMLName", result[2], "XMLName: { }")
	assertEqual(t, "CountryCode", result[3], "CountryCode: TEST")
	assertEqual(t, "VatNumber", result[4], "VatNumber: 1234567890")
	assertEqual(t, "RequestDate", result[5], "RequestDate: ")
	assertEqual(t, "Valid", result[6], "Valid: true")
	assertEqual(t, "Name", result[7], "Name: Test name")
	assertEqual(t, "Address", result[8], "Address: Test Address")
}
