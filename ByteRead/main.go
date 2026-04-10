package main

import (
    "fmt"
    "os"
)

func main() {
    filepath := "data.txt"

    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    data := make([]byte, 10)
    _, err = file.Read(data)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(data)
    fmt.Println(string(data))
}