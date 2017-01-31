package main

import (
    "os"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type ServerConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Request struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type Response struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

type Route struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Routes []Route

type Config struct {
	Server ServerConfig `json:"server"`
	Routes Routes       `json:"routes"`
}

type Server struct {
	Config Config
}

func (c Server) getResponse(r *http.Request) Response {
	for _, route := range c.Config.Routes {
		if route.Request.Url == r.URL.Path {
			if route.Request.Method == r.Method {
				return route.Response
			}
		}
	}

	return Response{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)}
}

func (c Server) Handle(w http.ResponseWriter, r *http.Request) {
	var response Response = c.getResponse(r)

	w.WriteHeader(response.Status)
	w.Write([]byte(response.Body))
}

func (c Server) Run() {
	fmt.Printf(
		"Listening at address %v on port %v\n",
		c.Config.Server.Address,
		c.Config.Server.Port,
	)
	var bind = fmt.Sprintf("%v:%v", c.Config.Server.Address, c.Config.Server.Port)
	http.HandleFunc("/", c.Handle)
	log.Fatal(http.ListenAndServe(bind, nil))
}

func main() {
    content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(content, &config)

	var server Server = Server{config}
	server.Run()
}
