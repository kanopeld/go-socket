## TCP Socket.io-like library

[![](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/kanopeld/go-socket)
[![Go Report Card](https://goreportcard.com/badge/github.com/kanopeld/go-socket)](https://goreportcard.com/report/github.com/kanopeld/go-socket)
![](https://github.com/kanopeld/go-socket/workflows/ci/badge.svg)

###Example usage you can see in example file into ./example directory.

###This library allows you to organize work with sockets through the convenient mechanism of "events". Here are some examples of use:

```go
s, err := server.NewServer(":6500")
if err != nil {
    panic(err)
}

err = s.On("connection", func(c core.SClient) {
    fmt.Println("Hello socket")
})
```

###In this example, we started listening and accepting connections on port 6500. As soon as the client connects to our system, the connection event will be triggered (it is basic and mandatory) in the callback of which the socket object will be transferred, with which we will work in the future
###In the example above, we considered the core.SClient interface that is used on the server side. There is also a core.DClient interface used on the client side. Their difference in the absence of broadcast sending at the client
