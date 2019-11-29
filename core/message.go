package core

var (
	CharStartEventName = []byte("[")[0]
	CharEndEventName   = []byte("]")[0]
)

type Message struct {
	EventName string
	Data      []byte
}

func (m Message) MarshalBinary() []byte {
	res := make([]byte, 0)
	res = append(res, CharStartEventName)
	res = append(res, []byte(m.EventName)...)
	res = append(res, CharEndEventName)
	res = append(res, m.Data...)
	return res
}
