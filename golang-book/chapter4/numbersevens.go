package main

import "fmt"

func main() {
    fmt.Println("Evens Numbers from 1 to 100:")
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i)
        }
    }
}
