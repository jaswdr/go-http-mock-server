package gomock

import "testing"

func BenchmarkCreateHttpMockServer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        CreateHttpMockServer("./example/mock.json")
    }
}

func TestCreateHttpMockServer(t *testing.T) {
    var server Server = CreateHttpMockServer("./example/mock.json")

    if server.Config.Server.Address != "0.0.0.0" {
        t.Errorf(
            "%s different that expected %s",
            "0.0.0.0",
            server.Config.Server.Address,
        )
    }

    if server.Config.Server.Port != 8080 {
        t.Errorf(
            "%s different that expected %s",
            8080,
            server.Config.Server.Port,
        )
    }

    var httpMethods []string = []string{"GET","POST","PUT","DELETE","HEAD","PATH"}

    for _, route := range server.Config.Routes {
        var validMethod = false
        for _, method := range httpMethods {
            if route.Request.Method == method {
                validMethod = true
            }
        }
        if validMethod == false {
            t.Errorf(
                "\"%s\" is not a valid HTTP method",
                route.Request.Method,
            )
        }

        if len(route.Request.Url) == 0 {
            t.Error(
                "Empty URL",
                route.Request.Url,
            )
        }

        if route.Response.Status <= 0 {
            t.Errorf(
                "\"%d\" is an invalid status",
                route.Response.Status,
            )
        }

        if len(route.Response.Body) == 0 {
            t.Errorf(
                "\"%s\" is an empty body",
                route.Response.Body,
            )
        }
    }
}
