package core

func DecodeMessage(data []byte) Message {
	var msg Message
	name := ""
	var start = false
	var end = false
	var endAt = 0
	for p, char := range data {
		if char == CharStartEventName {
			start = true
			continue
		} else if char == CharEndEventName {
			start = false
			end = true
		}
		if start {
			name += string(char)
		}
		if end {
			endAt = p
			break
		}
	}
	msg.EventName = name
	msg.Data = data[endAt+1:]
	return msg
}
