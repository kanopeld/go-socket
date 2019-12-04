package socket

import (
	"net"
	"reflect"
)

type Emitter interface {
	Emit(event string, arg interface{}) error
	sender
}

func getEmitter(c net.Conn) Emitter {
	return &defaultEmitter{getSender(c)}
}

type defaultEmitter struct {
	sender
}

func (de *defaultEmitter) Emit(event string, arg interface{}) error {
	var data []byte
	t := reflect.TypeOf(arg)
	if t != nil {
		switch t.Kind() {
		case reflect.Slice:
			tryData, ok := arg.([]byte)
			if !ok {
				return ErrUnsupportedArgType
			}
			data = tryData
		case reflect.String:
			data = []byte(arg.(string))
		default:
			return ErrUnsupportedArgType
		}
	}
	return de.send(&Package{PT: PackTypeEvent, Payload: Message{EventName: event, Data: data}.MarshalBinary()})
}
