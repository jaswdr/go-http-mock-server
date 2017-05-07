package main

import (
    "fmt"
    "github.com/jaschweder/gomock"
)

func main() {
    fmt.Println("Hello World")
    var server := gomock.CreateHttpMockServer("example/mock.json")
}
