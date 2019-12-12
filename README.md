# go-socket

TCP Socket.io-like library

[![](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/kanopeld/socket)
[![Go Report Card](https://goreportcard.com/badge/github.com/kanopeld/socket)](https://goreportcard.com/report/github.com/kanopeld/socket)
![](https://github.com/kanopeld/socket/workflows/ci/badge.svg)
![](https://coveralls.io/github/igm)

### Quick start

This library allows you to organize work with sockets through the convenient mechanism of "events". Here are some examples of use:

```go
package main

import (
	"fmt"
	"github.com/kanopeld/socket"
	"time"
)

func main() {
	// create new socket  server
    s, err := socket.NewServer(":6500")
    if err != nil {
        panic(err)
    }

    // When connecting a new client, the connection event will be raised, therefore for this to work, such a handler must be defined
    // ConnectionName=="connection"
    s.On(socket.ConnectionName, func(c socket.Client, data []byte) error {
        fmt.Println("connected", c.ID())

        // All other handlers we assign to the newly created socket
        c.On("test", func(c socket.Client, data []byte) error {
            fmt.Println("server got test event")
            fmt.Printf("Test (%s) message\n", string(data))
            _ = c.Emit("test", nil)
            return nil
        })

        _ = c.Broadcast("test1", nil)

        // DisconnectionName=="disconnection"
        c.On(socket.DisconnectionName, func(c socket.Client, data []byte) error {
            fmt.Println("Server disc")
            return nil
        })
        return nil
    })

    // this is a blocking method so we will spawn a goroutine for it. Use s.Stop() to stop the server
    go s.Start()

    // Thus, we establish a connection to the  At the time of opening, the server receives a message and a connection event is called on it
    d1, err := socket.NewDial("localhost:6500")
    if err != nil {
        panic(err)
    }
    d1.On(socket.ConnectionName, func(c socket.Client, data []byte) error {
        fmt.Println("d1 Connect!")
        c.On("test", func(c socket.Client, data []byte) error {
            go func() {
                fmt.Println("d1 got test event")
            }()
            return nil
        })
        _ = c.Emit("test", []byte("hello"))

        c.On(socket.DisconnectionName, func(c socket.Client, data []byte) error {
            fmt.Println("d1 disc")
            return nil
        })

        c.On("test1", func(c socket.Client, data []byte) error {
            fmt.Println("d1 got dial broadcast")
            return nil
        })
        return nil
    })

    d2, err := socket.NewDial("localhost:6500")
    if err != nil {
        panic(err)
    }
    d2.On(socket.ConnectionName, func(c socket.Client, data []byte) error {
        c.On("test", func(c socket.Client, data []byte) error {
            go func() {
                fmt.Println("d2 got test event")
            }()
            return nil
        })
        _ = c.Emit("test", []byte("hello"))

        c.On(socket.DisconnectionName, func(c socket.Client, data []byte) error {
            fmt.Println("d2 disc")
            return nil
        })

        c.On("test1", func(c socket.Client, data []byte) error {
            fmt.Println("d2 got dial 1 broadcast")
            return nil
        })
        return nil
    })

    // for make sure what dial code finished
    time.Sleep(10 * time.Second)

    // stop the server wait & close tcp connect
    s.Stop()
}
```

In this example, we start listening and accepting connections on port 6500. As soon as a client connects to our server, the 'connection' event is triggered on both ends. Inside of our 'connection' event we define the remaining events on the received client object. We also use a predefined 'disconnection' event that gets triggered upon disconnection.\
In the example above, we considered the SClient interface that is used on the server side. There is also a DClient interface used on the client side. Their difference in the absence of broadcast sending at the client

**For a more thorough documentation see [godoc](https://godoc.org/github.com/kanopeld/socket)**

The project uses a standard MIT license. See more in LICENSE
