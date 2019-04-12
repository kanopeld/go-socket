package socket

type Message struct {
	EventName string
	Data      []byte
}

func (m Message) MarshalBinary() []byte {
	res := make([]byte, 0)
	res = append(res, []byte("[")...)
	res = append(res, []byte(m.EventName)...)
	res = append(res, []byte("]")...)
	res = append(res, m.Data...)
	return res
}
