package socket

import (
	"testing"
	"log"
	"time"
	"bufio"
	"strings"
)

type TestData struct {
	Data string `json:"data"`
}

func TestNewMessageEvent(t *testing.T) {
	msg, err := NewEncodeMessage(_EVENT, "add", "Hello Socket")
	if err != nil {
		t.Fatal(err)
	}

	reader, err := bufio.NewReader(strings.NewReader(msg)).ReadBytes('\n')
	dec, err := NewMessageDecoder(reader)
	if err != nil {
		t.Error(err)
		return
	}

	if dec.payload != "Hello Socket" {
		t.Fatal("Payload invalid")
	}

	if dec.mt != _EVENT {
		t.Error("error parse message type")
		return
	}

	if dec.eventName != "inst" {
		t.Error("Error parse event name")
		return
	}

	log.Printf("dec: %v", dec)
}

func TestNewMessageDecoderConnection(t *testing.T) {
	str := "0"
	dec, err := NewMessageDecoder([]byte(str))
	if err != nil {
		t.Error(err)
		return
	}

	if dec.mt != _CONNECTION {
		t.Error("Error parse message type")
		return
	}
}

func TestMainFunc(t *testing.T) {
	server, err := NewServer(":24641")
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	server.On("connection", func(client Client, msg string) {
		client.On("inst", func(msg string) {
			log.Print(msg)
		})
	})

	dial, err := NewDial("127.0.0.1:24641")
	if err != nil {
		t.Fatal(err)
	}

	defer dial.Close()

	dial.Send("inst", "Hello")
	time.Sleep(20 * time.Second)
}