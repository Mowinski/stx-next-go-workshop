package main

import (
	"fmt"
	"os"

	"github.com/Mowinski/stx-next-go-workshop/vatno"
)

func abort(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	vatNo, err := vatno.Read(os.Stdin)
	if err != nil {
		abort(err)
	}
	fmt.Println(vatNo)
}
