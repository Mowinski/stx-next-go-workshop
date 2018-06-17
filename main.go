package main

import (
	"fmt"
	"github.com/Mowinski/stx-next-go-workshop/vies"
	"io"
	"os"
)

func readFrom(reader io.Reader) (data vies.InputData, err error) {
	fmt.Print("Enter country code: ")
	_, err = fmt.Fscanf(reader, "%s", &data.CountryCode)
	if err != nil {
		return data, err
	}

	fmt.Print("Enter VAT: ")
	_, err = fmt.Fscanf(reader, "%s", &data.VATNumber)
	if err != nil {
		return data, err
	}

	return data, nil
}

func writeTo(writer io.Writer, vatDetail *vies.VatDetails) {
	fmt.Fprintln(writer, "Error: ", vatDetail.Error)
	fmt.Fprintln(writer, "ErrorCode: ", vatDetail.ErrorCode)
	fmt.Fprintln(writer, "XMLName: ", vatDetail.XMLName)
	fmt.Fprintln(writer, "CountryCode: ", vatDetail.CountryCode)
	fmt.Fprintln(writer, "VatNumber: ", vatDetail.VatNumber)
	fmt.Fprintln(writer, "RequestDate: ", vatDetail.RequestDate)
	fmt.Fprintln(writer, "Valid: ", vatDetail.Valid)
	fmt.Fprintln(writer, "Name: ", vatDetail.Name)
	fmt.Fprintln(writer, "Address: ", vatDetail.Address)
}

func main() {
	inputData, err := readFrom(os.Stdin)
	if err != nil {
		panic(err)
	}

	vatDetail, err := vies.Check(inputData)
	if err != nil {
		panic(err)
	}

	writeTo(os.Stdout, vatDetail)
}
