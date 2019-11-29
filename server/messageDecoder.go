package server

func DecodeMessage(data []byte) Message {
	var msg Message
	name := ""
	var start = false
	var end = false
	var endAt = 0
	for p, char := range data {
		var ch = string(char)
		if ch == "[" {
			start = true
			continue
		} else if ch == "]" {
			start = false
			end = true
		}

		if start {
			name += ch
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
