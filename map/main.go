package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"greee": "#4bf754",
	// }

	colors["white"] = "#ffffff"

	delete(colors, "white") // delete a key-value pair in map

	fmt.Println(colors)
}
