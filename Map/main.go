package main

import (
	"fmt"
)

/*
func main() {
	//var colors map[string]string // 1st

	//colors := make(map[string]string)//2nd

	colors := make(map[int]string)

	//colors["white"] = "#ffffff"

	colors[10] = "#ffffff"
	delete(colors, 10)

	fmt.Println(colors)
}
*/

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
