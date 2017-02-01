package main

import (
    "os"
    "fmt"
    "github.com/jaschweder/gomock"
)

func main() {

    if len(os.Args) < 2 {
        fmt.Println(`
Go HTTP Mock Server

Usage:
    gomock <config_file>
        `)
        os.Exit(1)
    }

    var server gomock.Server = gomock.CreateHttpMockServer(os.Args[1])

    server.Run()
}
