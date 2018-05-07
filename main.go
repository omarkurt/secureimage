package main

import (
	"fmt"
	"github.com/c1982/secureimage"
)

func main() {
	trusted, err := secureimage.Check("testdata/1N3.jpeg")

	if err != nil {
		panic(err)
	}

	if trusted {
		fmt.Println("file is trusted.") // Filename ?
	} else {
		fmt.Println("bad file")
	}
}

// output
//file is trusted.
