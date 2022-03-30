package main

import (
	// "fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, file)

	}

}
