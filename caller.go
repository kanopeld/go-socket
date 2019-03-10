package socket

import (
	"fmt"
	"reflect"
)

type caller struct {
	Func       reflect.Value
	Args       []reflect.Type
	NeedSocket bool
}

func NewCaller(f interface{}) (*caller, error) {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		return nil, fmt.Errorf("f is not func")
	}
	ft := fv.Type()
	if ft.NumIn() == 0 {
		return &caller{
			Func: fv,
		}, nil
	}
	args := make([]reflect.Type, ft.NumIn())
	for i, n := 0, ft.NumIn(); i < n; i++ {
		args[i] = ft.In(i)
	}
	needSocket := false
	if args[0].Name() == "Client" {
		args = args[1:]
		needSocket = true
	}
	return &caller{
		Func:       fv,
		Args:       args,
		NeedSocket: needSocket,
	}, nil
}

func (c *caller) GetArgs() []interface{} {
	ret := make([]interface{}, len(c.Args))
	for i, argT := range c.Args {
		if argT.Kind() == reflect.Ptr {
			argT = argT.Elem()
		}
		v := reflect.New(argT)
		ret[i] = v.Interface()
	}
	return ret
}

func (c *caller) Call(so Client, data []byte) []reflect.Value {
	a := make([]reflect.Value, 0)
	if c.NeedSocket {
		a = append(a, reflect.ValueOf(so))
	}

	if len(c.Args) > 0 {
		a = append(a, reflect.ValueOf(data))
	}

	return c.Func.Call(a)
}
