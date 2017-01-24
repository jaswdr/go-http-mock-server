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
		"port": 8080,
		"urlPrefix": "/api"
	},
	"record": {
		"saveTo": "/tmp/gomock",
		"notifyConsole:": true
	},
	"routes": [{
		"request": {
			"method": "GET",
			"url": "/v1/hello"
		},
		"response": {
			"status": 200,
			"body": "Hello World!"
		}
	}]
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
