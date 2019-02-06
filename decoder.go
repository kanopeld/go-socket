package socket

import (
	"strconv"
	"errors"
)

const (
	_CONNECTION = iota
	_ASK
	_EVENT
	_DISCCONNECTION
	_ERROR
)

const (
	CONNECTION_NAME = "connection"
	DISCONNECTION_NAME = "disconnection"
)

const (
	MAX_EVENT_LENGTH = 999
)

var (
	ErrDecodeFlag = errors.New("error decode flag")
	ErrNotFoundEventAction = errors.New("not found event action")
	ErrEmptyMessage = errors.New("empty message")
)

type Decoder struct {
	mt int
	eventName string
	payload string
}

func NewMessageDecoder(msg []byte) (*Decoder, error) {
	if len(msg) <= 1 {
		return nil, ErrEmptyMessage
	}

	dec := &Decoder{}
	mt := parseFlag(msg[:1])
	if mt == -1 {
		return nil, ErrDecodeFlag
	}
	msg = msg[1:]
	dec.mt = mt
	switch mt {
	case _CONNECTION:
		return dec, nil
	case _EVENT:
		nameLength := parseFlag(msg[:3])
		if nameLength == -1 {
			return nil, ErrDecodeFlag
		}
		msg = msg[3:]

		eventName := msg[:nameLength]
		msg = msg[nameLength:]
		dec.eventName = string(eventName)
		dec.payload = string(msg[:len(msg) - 1])
		return dec, nil
	case _ERROR:
		return dec, nil
	case _DISCCONNECTION:
		return dec, nil
	}

	return nil, ErrNotFoundEventAction
}

func parseFlag(b []byte) int {
	stF := string(b)
	mt, err := strconv.Atoi(stF)
	if err != nil {
		return -1
	}
	return mt
}
