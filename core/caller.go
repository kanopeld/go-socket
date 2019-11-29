package core

import "reflect"

type Caller interface {
	Call(so Client, data []byte) []reflect.Value
}
