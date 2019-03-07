package socket

import (
	"encoding/json"
	"testing"
	"log"
	"time"
)

func TestNewServer(t *testing.T) {
	s, err := NewServer(":6500")
	if err != nil {
		t.Fatal(err)
	}

	err = s.On("connection", func(c Client) {
		log.Printf("connected %s", c.ID())

		err = c.On("test", func(data []byte) {
			var st string
			err := json.Unmarshal(data, &st)
			if err != nil {
				t.Fatal(err)
			}
			log.Printf("%s", st)
		})

		if err != nil {
			t.Fatal(err)
		}
	})

	if err != nil {
		t.Fatal(err)
	}

	go s.Start()

	d, err := NewDial("localhost:6500")
	if err != nil {
		t.Fatal(err)
	}
	err = d.On("connection", func(c Client) {
		b, err := json.Marshal("test")
		if err != nil {
			t.Fatal(err)
		}
		if err := d.Emit("test", b); err != nil {
			t.Fatal(err)
		}
	})

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	s.Stop()
}
