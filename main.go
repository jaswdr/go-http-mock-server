package main

import (
	"fmt"
  "io"
	"io/ioutil"
	"log"
  "encoding/json"
  "net/http"
)

const configFile string = "gomock.json"

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

func LoadConfig() Config {
    content, err := ioutil.ReadFile(configFile)

    if err != nil {
        log.Fatal(err)
    }

    var config Config

    json.Unmarshal(content, &config)

    return config
}

type Server struct {
  Config Config
}

func (c Server) Handler(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "hello world!\n")
}

func (c Server) Run() {
  var bind = fmt.Sprintf(":%v", c.Config.Server.Port);
  log.Fatal( http.ListenAndServe(bind, nil))
}

func main() {
    var config Config = LoadConfig()
    var server Server = Server{config}
    server.Run()
}
