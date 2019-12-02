package main

import (
	"fmt"
	"github.com/kanopeld/go-socket/core"
	"github.com/kanopeld/go-socket/dial"
	"github.com/kanopeld/go-socket/server"
	"time"
)

func main() {
	//create new server
	s, err := server.NewServer(":6500")
	if err != nil {
		panic(err)
	}

	//when new client connecting, server will call "connection" event.
	err = s.On("connection", func(c core.SClient) {
		fmt.Printf("connected %s\n", c.ID())

		err = c.On("test", func(c core.SClient, data []byte) {
			fmt.Println("server got test event")
			fmt.Printf("Test (%s) message\n", string(data))
			_ = c.Emit("test", nil)
		})

		if err != nil {
			panic(err)
		}

		_ = c.Broadcast("test1", nil)

		_ = c.On("disconnection", func() {
			fmt.Println("Server disc")
		})
	})

	if err != nil {
		panic(err)
	}

	//this method will block next code and wait when program finish or will called Stop() method
	go s.Start()

	d, err := dial.NewDial("localhost:6500")
	if err != nil {
		panic(err)
	}
	err = d.On("connection", func(c core.DClient) {
		fmt.Print("dial Commect!")
		_ = c.On("test", func() {
			go func() {
				fmt.Println("dial got test event")
			}()
		})
		_ = c.Emit("test", "hello")

		_ = c.On("disconnection", func() {
			fmt.Println("Dial disc")
		})

		_ = c.On("test1", func() {
			fmt.Printf("got dial broadcast")
		})
	})

	d1, err := dial.NewDial("localhost:6500")
	if err != nil {
		panic(err)
	}
	err = d1.On("connection", func(c core.DClient) {
		_ = c.On("test", func() {
			go func() {
				fmt.Println("dial got test event")
			}()
		})
		_ = c.Emit("test", "hello")

		_ = c.On("disconnection", func() {
			fmt.Println("Dial disc")
		})

		_ = c.On("test1", func() {
			fmt.Printf("got dial 1 broadcast")
		})
	})

	//for make sure what dial code finished
	time.Sleep(10 * time.Second)

	//stop the server wait & close tcp connect
	s.Stop()
}
