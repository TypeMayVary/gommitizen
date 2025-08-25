package main

import (
    "fmt"
    "gommitizen/pkg/utils"
)

func main() {
    fmt.Println("Welcome to my Go project!")
    // Example usage of utility functions
    str := "123"
    num, err := utils.StringToInt(str)
    if err != nil {
        fmt.Printf("Error converting string to int: %v\n", err)
        return
    }
    fmt.Printf("Converted string to int: %d\n", num)
    strBack := utils.IntToString(num)
    fmt.Printf("Converted int back to string: %s\n", strBack)
}