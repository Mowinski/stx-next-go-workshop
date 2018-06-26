package main

import (
	"fmt"
	"os"

	"github.com/Mowinski/stx-next-go-workshop/vatno"
	"github.com/Mowinski/stx-next-go-workshop/vies"
)

func abort(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	for {
		vatNo, err := vatno.Read(os.Stdin)
		if err != nil {
			abort(err)
		}

		vatDetail, err := vies.Query(vatNo)
		if err != nil {
			abort(err)
		}

		vatDetail.Print(os.Stdout)
	}
}
