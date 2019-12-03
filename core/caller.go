package core

import (
	"errors"
	"reflect"
)

//The interface
type Caller interface {
	Call(so Client, data []byte) []reflect.Value
	GetArgs() []interface{}
	Socket() bool
}

var (
	//ErrUnsupportedArgType This error can be received in case of transfer in callback of not supported argument.
	ErrUnsupportedArgType   = errors.New("received arg type is not support")
	ErrTooManyArgsForCaller = errors.New("error too many argument for caller func")
	//ErrFIsNotFunc The argument for callback must be a function.
	ErrFIsNotFunc = errors.New("f ins not a func")
)

type caller struct {
	Func       reflect.Value
	Args       []reflect.Type
	NeedSocket bool
}

//GetCaller function return a Caller interface. The parameter "name"
//is used in the closure of the function and in the future to check
//which interface (its name) of the client is passed to it (function).
//If this data is different, there will be an error
func GetCaller(name string) func(f interface{}) (Caller, error) {
	return func(f interface{}) (c Caller, err error) {
		fv := reflect.ValueOf(f)
		if fv.Kind() != reflect.Func {
			return nil, ErrFIsNotFunc
		}
		ft := fv.Type()
		if ft.NumIn() == 0 {
			return &caller{
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
				if v.Name() != name {
					return nil, ErrUnsupportedArgType
				}
			default:
				return nil, ErrUnsupportedArgType
			}
			if v.Name() == name && i == 0 {
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
		return &caller{
			Func:       fv,
			Args:       args,
			NeedSocket: needSocket,
		}, nil
	}
}

//GetArgs function returns a slice of arguments. Used mainly in tests
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

//Call function calls the callback passed at creation. May panic if received arguments wrong
func (c *caller) Call(so Client, data []byte) []reflect.Value {
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

//Socket function returns the flag whether it needs a socket object to run.
//The flag is set at the stage of Caller creation.
func (c *caller) Socket() bool {
	return c.NeedSocket
}
