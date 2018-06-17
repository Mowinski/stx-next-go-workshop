package main

import (
	"fmt"

	"github.com/Mowinski/stx-next-go-workshop/vies"
)

func readFromStdIn() (data vies.InputData, err error) {
	fmt.Print("Enter country code: ")
	_, err = fmt.Scanf("%s", &data.CountryCode)
	if err != nil {
		return data, err
	}

	fmt.Print("Enter VAT: ")
	_, err = fmt.Scanf("%s", &data.VATNumber)
	if err != nil {
		return data, err
	}

	return data, nil
}

func writeToStdOut(vatDetail *vies.VatDetails) {
	fmt.Println("Error: ", vatDetail.Error)
	fmt.Println("ErrorCode: ", vatDetail.ErrorCode)
	fmt.Println("XMLName: ", vatDetail.XMLName)
	fmt.Println("CountryCode: ", vatDetail.CountryCode)
	fmt.Println("VatNumber: ", vatDetail.VatNumber)
	fmt.Println("RequestDate: ", vatDetail.RequestDate)
	fmt.Println("Valid: ", vatDetail.Valid)
	fmt.Println("Name: ", vatDetail.Name)
	fmt.Println("Address: ", vatDetail.Address)
}

func main() {
	inputData, err := readFromStdIn()
	if err != nil {
		panic(err)
	}

	vatDetail, err := vies.Check(inputData)
	if err != nil {
		panic(err)
	}

	writeToStdOut(vatDetail)
}
