package main

import (
	"fmt"
)

func main() {

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// var colors map[string]string //we use this way to use map when we want to call out the properties later in time

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf475",
		"white": "#ffffff",
	}

	// delete(colors, "white") //when you want to delete an item from a map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, ": ", hex)
	}
}
