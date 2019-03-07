package socket

import "encoding/json"

type Message struct {
	EventName string
	Data []byte
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
