package socket

import "errors"

var (
	ErrTooManyArgsForCaller = errors.New("error too many argument for caller func")
	//ErrFIsNotFunc The argument for callback must be a function.
	ErrFIsNotFunc = errors.New("f ins not a func")
)
