package server

import (
	"errors"
	"fmt"
	"github.com/kanopeld/go-socket/core"
	"reflect"
)

var (
	ErrUnsupportedArgType   = errors.New("received arg type is not support")
	ErrTooManyArgsForCaller = errors.New("error too many argument for caller func")
)

type serverCaller struct {
	Func       reflect.Value
	Args       []reflect.Type
	NeedSocket bool
}

func NewCaller(f interface{}) (core.Caller, error) {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		return nil, fmt.Errorf("f is not func")
	}
	ft := fv.Type()
	if ft.NumIn() == 0 {
		return &serverCaller{
			Func: fv,
		}, nil
	}
	needSocket := false
	args := make([]reflect.Type, 0)
	for i, n := 0, ft.NumIn(); i < n; i++ {
		v := ft.In(i)
		switch v.Kind() {
		case reflect.String:
		case reflect.Slice:
		case reflect.Interface:
			if v.Name() != "Client" {
				return nil, ErrUnsupportedArgType
			}
		default:
			return nil, ErrUnsupportedArgType
		}
		if v.Name() == "Client" && i == 0 {
			needSocket = true
		}
		if needSocket && i >= 2 {
			return nil, ErrTooManyArgsForCaller
		} else if !needSocket && i >= 1 {
			return nil, ErrTooManyArgsForCaller
		}
		args = append(args, ft.In(i))
	}
	if needSocket {
		args = args[1:]
	}
	return &serverCaller{
		Func:       fv,
		Args:       args,
		NeedSocket: needSocket,
	}, nil
}

func (c *serverCaller) GetArgs() []interface{} {
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

func (c *serverCaller) Call(so core.Client, data []byte) []reflect.Value {
	a := make([]reflect.Value, 0)
	if c.NeedSocket {
		a = append(a, reflect.ValueOf(so))
	}
	if len(c.Args) > 0 {
		if c.Args[len(c.Args)-1].Kind() == reflect.String {
			a = append(a, reflect.ValueOf(string(data)))
		} else {
			a = append(a, reflect.ValueOf(data))
		}
	}
	return c.Func.Call(a)
}
