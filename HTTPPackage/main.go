package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	getResonse()
}

func getResonse() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {

		io.Copy(os.Stdout, resp.Body)
	}
}
