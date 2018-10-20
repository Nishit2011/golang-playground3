package main

import "fmt"

func main() {
	//[string] refers to type of key while string refers to type of value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
