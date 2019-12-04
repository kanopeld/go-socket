package socket

var (
	charStartEventName = []byte("[")[0]
	charEndEventName   = []byte("]")[0]
)

// Message stores information about a single message
type Message struct {
	EventName string
	Data      []byte
}

// MarshalBinary serializes a message into bytes
func (m Message) MarshalBinary() []byte {
	res := make([]byte, 0)
	res = append(res, charStartEventName)
	res = append(res, []byte(m.EventName)...)
	res = append(res, charEndEventName)
	res = append(res, m.Data...)
	return res
}
