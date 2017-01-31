# Go http mock server
> Mock your APIs in a really fast, robust and comprehensive way

## Install
`go get github.com/jaschweder/go-http-mock-server`

## Getting start

Create a ``mock.json`` file like this:
```json
{
    "server":{
        "address":"0.0.0.0",
        "port":8080
    },
    "routes":[
        {
            "request":{
                "method":"GET",
                "url":"/hello"
            },
            "response":{
                "status":200,
                "body":"{\"message\": \"World!\"}"
            }
        }
    ]
}
```

Run ``gomock``:
```shell
$ gomock mock.json
```

Try make a request:
```shell
$ curl http://localhost:8080/hello
{message: "World!"}
```

That's it! You rock!!!

## How to contribute:

Issue report and PR's are welcome

## Author:
Created by Jonathan A. Schweder <jonathanschweder@gmail.com>

## License
MIT
