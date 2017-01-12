package main

import (
	"fmt"
	"io/ioutil"
	"log"
    "encoding/json"
)

type ServerConfig struct {
	Port      int    `json:"port"`
	UrlPrefix string `json:"urlPrefix"`
}

type RecordConfig struct {
	SaveTo        string `json:"saveTo"`
	NotifyConsole bool   `json:"notifyConsole"`
}

type Request struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

type Response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

type Route struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Routes []Route

type Config struct {
	Server ServerConfig `json:"server"`
	Record RecordConfig `json:"record"`
	Routes Routes       `json:"routes"`
}

func ReadConfig() Config {
    content, err := ioutil.ReadFile("gomock.json")

    if err != nil {
        log.Fatal(err)
    }

    var config Config

    json.Unmarshal(content, &config)

    return config
}

func main() {
    var config Config = ReadConfig()

    fmt.Printf("%s", config)
}
