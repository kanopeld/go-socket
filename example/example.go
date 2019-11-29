package main

import (
	"fmt"
	"github.com/kanopeld/go-socket"
	"time"
)

func main() {
	//create new server
	s, err := socket.NewServer(":6500")
	if err != nil {
		panic(err)
	}

	//when new client connecting, server will call "connection" event.
	err = s.On(socket.CONNECTION_NAME, func(c socket.Client) {
		fmt.Printf("connected %s\n", c.ID())

		err = c.On("test", func(c socket.Client, data []byte) {
			fmt.Println("server got test event")
			fmt.Printf("Test (%s) message\n", string(data))
			_ = c.Emit("test", nil)
		})

		if err != nil {
			panic(err)
		}

		_ = c.Broadcast("test1", nil)

		_ = c.On(socket.DISCONNECTION_NAME, func() {
			fmt.Println("Server disc")
		})
	})

	if err != nil {
		panic(err)
	}

	//this method will block next code and wait when program finish or will called Stop() method
	go s.Start()

	d, err := socket.NewDial("localhost:6500")
	if err != nil {
		panic(err)
	}
	err = d.On(socket.CONNECTION_NAME, func(c socket.Client) {
		_ = d.On("test", func() {
			go func() {
				fmt.Println("dial got test event")
			}()
		})
		_ = d.Emit("test", "hello")

		_ = d.On(socket.DISCONNECTION_NAME, func() {
			fmt.Println("Dial disc")
		})

		_ = d.On("test1", func() {
			fmt.Printf("got dial broadcast")
		})
	})

	d1, err := socket.NewDial("localhost:6500")
	if err != nil {
		panic(err)
	}
	err = d1.On(socket.CONNECTION_NAME, func(c socket.Client) {
		_ = d1.On("test", func() {
			go func() {
				fmt.Println("dial got test event")
			}()
		})
		_ = d1.Emit("test", "hello")

		_ = d1.On(socket.DISCONNECTION_NAME, func() {
			fmt.Println("Dial disc")
		})

		_ = d1.On("test1", func() {
			fmt.Printf("got dial 1 broadcast")
		})
	})

	//for make sure what dial code finished
	time.Sleep(5 * time.Second)

	//stop the server wait & close tcp connect
	s.Stop()
}
