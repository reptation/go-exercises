package main

import (
    "fmt"
)

func main() {
    var i int
    fmt.Print("Type a number: ")
    fmt.Scan(&i)

    if (i % 3 == 0) && (i % 5 == 0) {
	fmt.Println("FizzBuzz")
    } else if i % 3 == 0 {
        fmt.Println("Fizz")
    } else if i % 5 == 0 {
        fmt.Println("Buzz")
    }
}
