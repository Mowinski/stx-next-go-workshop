package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
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

	if result.CountryCode != "TEST" {
		t.Errorf("Wrong CountryCode, expected 'TEST', got '%s'", result.CountryCode)
	}

	if result.VATNumber != "1234567890" {
		t.Errorf("Wrong VATNUmber, expected '1234567890', got '%s'", result.VATNumber)
	}
}
