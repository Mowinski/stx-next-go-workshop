package vies

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

type InputData struct {
	CountryCode string
	VATNumber   string
}

type VatDetails struct {
	XMLName     xml.Name
	CountryCode string
	VatNumber   string
	RequestDate string
	Valid       bool
	Name        string
	Address     string
	Error       string
	ErrorCode   string
}

const viesReqPayloadTmpl = ``

const viesServiceUrl = ""

func prepareViesRequestBody(vatDetail InputData) *bytes.Buffer {
	return nil
}

func sendViesPostRequest(body *bytes.Buffer) (*http.Response, error) {
	return nil, nil
}

func encodeResponseFromVies(body io.Reader) (vatDetails VatDetails, err error) {
	return vatDetails, err
}

func Check(vatDetail InputData) (*VatDetails, error) {
	return nil, nil
}
