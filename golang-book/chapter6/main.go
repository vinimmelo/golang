package main

import (
	"fmt"
)

func main() {
	elements := map[string]string{
		"He": "Hydrogen",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	if el, ok := elements["O"]; ok {
		fmt.Println(el)
	}
	fmt.Println("End of the program")
}
