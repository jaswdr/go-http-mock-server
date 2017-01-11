package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	Request  Request
	Response Response
}

type Routes []Route

type Config struct {
	Server ServerConfig
	Record RecordConfig
	Routes Routes
}

func main() {
	configFileContent, err := ioutil.ReadFile("gomock.json")

	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(configFileContent, config)

	fmt.Print(config)
}
