package socket

type EventCallback func(client Client, msg string)

type Event struct {
	name string
	callback EventCallback
}

func NewEvent(name string, callback EventCallback) *Event {
	return &Event{name:name,callback:callback}
}

type ClientEventCallback func(msg string)

type ClientEvent struct {
	name string
	callback ClientEventCallback
}

func NewClientEvent(name string, callback ClientEventCallback) *ClientEvent {
	return  &ClientEvent{name:name,callback:callback}
}