package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ServerConfig struct {
	Address   string    `json:"address"`
	Port      int    `json:"port"`
}

type RecordConfig struct {
	SaveTo        string `json:"saveTo"`
	NotifyConsole bool   `json:"notifyConsole"`
}

type Request struct {
	Url    string `json:"url"`
	Method string `json:"method"`
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

type Server struct {
	Config Config
}

func (c Server) Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!\n")
}

func (c Server) Run() {
    fmt.Printf(
        "Listening at address %v on port %v\n",
        c.Config.Server.Address,
        c.Config.Server.Port,
    )
    var bind = fmt.Sprintf("%v:%v", c.Config.Server.Address, c.Config.Server.Port)
    log.Fatal(http.ListenAndServe(bind, nil))
}

func main() {
	content, err := ioutil.ReadFile("gomock.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(content, &config)

	var server Server = Server{config}
	server.Run()
}
