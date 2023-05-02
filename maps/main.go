package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}
	colors["white"] = "#ffffff"
	printMap(colors)

}

func printMap(c map[string]string) {
	for colors, hex := range c {
		fmt.Printf("Hex code for %s is %s\n", colors, hex)
	}
}
