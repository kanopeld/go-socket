package socket

import "encoding/json"

type MessagePayload struct {
	Data []byte
}

type Message struct {
	EventName string
	Data MessagePayload
}

func (m *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (mp *MessagePayload) String() (string) {
	return string(mp.Data)
}