package socket

// Message stores information about a single message
type Message struct {
	EventName string
	Data      []byte
}

// MarshalBinary serializes a message into bytes
func (m Message) MarshalBinary() []byte {
	res := make([]byte, 0)
	res = append(res, []byte("[")...)
	res = append(res, []byte(m.EventName)...)
	res = append(res, []byte("]")...)
	res = append(res, m.Data...)
	return res
}
