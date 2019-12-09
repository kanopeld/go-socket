package socket

const (
	charStartEventName = byte('[')
	charEndEventName   = byte(']')
)

// Message stores information about a single message
type Message struct {
	// EventName is a custom event name chosen by the user
	EventName string
	// Data stores all the sent data
	Data []byte
}

// MarshalBinary serializes a message into bytes
func (m Message) MarshalBinary() (res []byte) {
	res = append(res, charStartEventName)
	res = append(res, []byte(m.EventName)...)
	res = append(res, charEndEventName)
	res = append(res, m.Data...)
	return
}

// decodeMessage splits a given message into its EventName and Data
func decodeMessage(data []byte) (msg Message) {
	start := false
	end := false
	endAt := 0
	for p, char := range data {
		if char == charStartEventName {
			start = true
			continue
		} else if char == charEndEventName {
			start = false
			end = true
		}
		if start {
			msg.EventName += string(char)
		}
		if end {
			endAt = p
			break
		}
	}
	msg.Data = data[endAt+1:]
	return
}
