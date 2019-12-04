package socket

import (
	"fmt"
	"github.com/kanopeld/socket"
	"time"
)

func Example_quickstart() {
	// create new socket  server
	s, err := socket.NewServer(":6500")
	if err != nil {
		panic(err)
	}

	// When connecting a new client, the connection event will be raised, therefore for this to work, such a handler must be defined
	// ConnectionName=="connection"
	err = s.On(socket.ConnectionName, func(c socket.SClient) {
		fmt.Println("connected", c.ID())

		// All other handlers we assign to the newly created socket
		err = c.On("test", func(c socket.SClient, data []byte) {
			fmt.Println("server got test event")
			fmt.Printf("Test (%s) message\n", string(data))
			_ = c.Emit("test", nil)
		})

		if err != nil {
			panic(err)
		}

		_ = c.Broadcast("test1", nil)

		// DisconnectionName=="disconnection"
		_ = c.On(socket.DisconnectionName, func() {
			fmt.Println("Server disc")
		})
	})

	if err != nil {
		panic(err)
	}

	// this is a blocking method so we will spawn a goroutine for it. Use s.Stop() to stop the server
	go s.Start()

	// Thus, we establish a connection to the  At the time of opening, the server receives a message and a connection event is called on it
	d1, err := socket.NewDial("localhost:6500")
	if err != nil {
		panic(err)
	}
	err = d1.On(socket.ConnectionName, func(c socket.DClient) {
		fmt.Println("d1 Connect!")
		_ = c.On("test", func() {
			go func() {
				fmt.Println("d1 got test event")
			}()
		})
		_ = c.Emit("test", "hello")

		_ = c.On(socket.DisconnectionName, func() {
			fmt.Println("d1 disc")
		})

		_ = c.On("test1", func() {
			fmt.Println("d1 got dial broadcast")
		})
	})

	d2, err := socket.NewDial("localhost:6500")
	if err != nil {
		panic(err)
	}
	err = d2.On(socket.ConnectionName, func(c socket.DClient) {
		_ = c.On("test", func() {
			go func() {
				fmt.Println("d2 got test event")
			}()
		})
		_ = c.Emit("test", "hello")

		_ = c.On(socket.DisconnectionName, func() {
			fmt.Println("d2 disc")
		})

		_ = c.On("test1", func() {
			fmt.Println("d2 got dial 1 broadcast")
		})
	})

	// for make sure what dial code finished
	time.Sleep(10 * time.Second)

	// stop the server wait & close tcp connect
	s.Stop()
}

func ExampleNewDial_dialling() {
	c, err := socket.NewDial("localhost:6500")
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Connection())

	// Outputs: net.Conn object
}

func ExampleNewServer_creation() {
	s, err := socket.NewServer(":6500")
	if err != nil {
		panic(err)
	}

	s.On(socket.ConnectionName, func(c socket.SClient) {
		// ...
	})

	go s.Start()
}
