package socket

import "encoding/json"

type Message struct {
	EventName string `json:"e"`
	Data []byte `json:"d"`
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
