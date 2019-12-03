package core

var (
	CharStartEventName = []byte("[")[0]
	CharEndEventName   = []byte("]")[0]
)

// Message stores information about a single message
type Message struct {
	EventName string
	Data      []byte
}

// MarshalBinary serializes a message into bytes
func (m Message) MarshalBinary() []byte {
	res := make([]byte, 0)
	res = append(res, CharStartEventName)
	res = append(res, []byte(m.EventName)...)
	res = append(res, CharEndEventName)
	res = append(res, m.Data...)
	return res
}
