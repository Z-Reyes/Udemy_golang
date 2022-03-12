package main

import "fmt"

type colorTable map[string]string

func main() {
	//var colors map[string]string
	//Equivalent to colors := make(map[string]string)

	//colors := make(map[string]string)
	colors := make(colorTable)
	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"
	colors["blue"] = "#0000ff"
	fmt.Println(colors["white"])
	colors.printMap()

}

func (c colorTable) printMap() {
	for i, hex := range c {
		fmt.Println("Hex code for", i, "is", hex)
	}
}
