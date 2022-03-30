package main

import (
	"fmt"
)

type englishBot struct{}
type spanishbot struct{}

type bot interface {
	getReply() string
}

func main() {

	eb := englishBot{}
	sb := spanishbot{}

	printReply(eb)
	printReply(sb)

}

func printReply(b bot) {
	fmt.Println(b.getReply())
}

func (englishBot) getReply() string {
	return "Hi there"
}

func (spanishbot) getReply() string {
	return "Hola Amigo"
}
