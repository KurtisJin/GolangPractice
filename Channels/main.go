package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //This is called "Blocking call". since we do not care for to get respons back, we leave it as _, err
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		os.Exit(1)
	} else {
		fmt.Println(link, "is up!")
		c <- "It's up!"
	}
}
