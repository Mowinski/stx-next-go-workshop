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

}

func main() {
	inputData, err := readFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(inputData)
}
