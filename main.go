package main

import (
	"fmt"
	"os"
	"github.com/c1982/secureimage"
)

func main() {
	trusted, err := secureimage.Check("./uploads/tmp_test.jpg")

	if err != nil {
		panic(err)
	}

	if trusted {
		fmt.Println("file is trusted.")
	} else {
		fmt.Println("bad file")
	}
}