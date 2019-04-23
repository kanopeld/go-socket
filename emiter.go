package socket

import (
	"errors"
	"net"
	"reflect"
)

var (
	ErrUnsupportedArgType = errors.New("received arg type is not support")
)

type defaultEmitter struct {
	c net.Conn
}

func (de *defaultEmitter) send(p *Package) error {
	_, err := de.c.Write(p.MarshalBinary())
	return err
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
	return de.send(&Package{PT: _PACKET_TYPE_EVENT, Payload: Message{EventName: event, Data: data}.MarshalBinary()})
}
