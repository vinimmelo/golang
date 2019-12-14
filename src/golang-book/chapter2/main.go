package main

import (
	"fmt"
	"os"
)

// this is a comment

func main() {
	// var name string
	// fmt.Print("Digite o seu nome: ")
	// fmt.Scanln(&name)
	// fmt.Println("Olá, seu nome é:", name)
    var x string = "Hello World"
    fmt.Println(x)
    x = "first "
    fmt.Println(x)
    x = x + "second"
    fmt.Println(x)
    y := 5
    fmt.Println(y)
	os.Exit(4)
}
