package socket

import (
	"strconv"
	"errors"
)

var (
	ErrEmptyName = errors.New("name is empty")
	ErrEventNameTooLong = errors.New("event name is too long")
)

func NewEncodeMessage(t int, name string, data string) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}

	msg := ""
	tStr := strconv.Itoa(t)
	msg = msg + tStr

	if t == _CONNECTION || t == _DISCCONNECTION {
		msg = msg + "\n"
		return msg, nil
	}

	nameL := CountStr(name)
	if nameL > MAX_EVENT_LENGTH {
		return "", ErrEventNameTooLong
	}

	nameLStr := strconv.Itoa(nameL)
	if nameL < 10 {
		nameLStr = "00"+nameLStr
	} else if nameL < 100 {
		nameLStr = "0"+nameLStr
	}

	msg = msg + nameLStr
	msg = msg + name

	if data == "" {
		msg = msg + "\n"
		return msg, nil
	}

	msg = msg + data
	msg = msg + "\n"
	return msg, nil
}
