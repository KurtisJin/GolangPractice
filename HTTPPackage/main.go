package main

import (
	"fmt"
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

		fmt.Println(resp.Body)
	}
}
