# GOlang mock server for testing
> Mock your APIs in a really fast, robust and comprehensive way

### Features:
 - Simple JSON file configuration
 - Powerful Request Matching
 - Record and Playback

Getting start
---

Create a ``gomock.json`` file like this:
```json
{
  "server": {
    "port": 8080,              # port bindded
    "urlPrefix": "/api"        # prefixed path of all routes
  },
  "record": {
    "saveTo": "/tmp/gomock",   # location where to save recorded request's
    "notifyConsole:": true     # indicate if print to stdout each time received a request
  },
  "routes": [
    {
      "request": {
        "method": "GET",       # method to match for this route
        "url": "/v1/hello"     # url to match for this route
      },
      "response": {
        "status": 200,         # response status
        "body": "Hello World!" # response body
      }
  ]
}
```

Run ``gomock`` passing the file path as parameter:
```shell
$ gomock ./gomock.json
```

Try make a request:
```shell
$ curl http://localhost:8080/api/v1/hello
Hello World!
```

Try replay a request:
```shell
$ gomock --replay ./gomock.json
```

That's it! You rock!!!

How to contribute:
---

Issue report and PR's are welcome

Author:
---
Created by Jonathan A. Schweder <jonathanschweder@gmail.com>
