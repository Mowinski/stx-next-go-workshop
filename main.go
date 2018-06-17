package main

import (
	"fmt"
	"github.com/Mowinski/stx-next-go-workshop/vies"
	"io"
	"os"
)

func readFrom(reader io.Reader) (data vies.InputData, err error) {
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
