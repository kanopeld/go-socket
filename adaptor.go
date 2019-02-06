package socket

import "sync"

type Adaptor interface {
	Call(name string, data string, c *Client)
	On(name string, callback EventCallback)
	GetEvent(name string) *Event
}

type DefaultAdaptor struct {
	el map[string]*Event
	sync.Mutex
}

func (da *DefaultAdaptor) Call(name string, data string, c *Client) {
	e := da.GetEvent(name)
	if e == nil {
		return
	}

	e.callback(*c, data)
}

func (da *DefaultAdaptor) GetEvent(name string) *Event {
	da.Lock()
	event, ok := da.el[name]
	da.Unlock()

	if !ok {
		return nil
	}

	return event
}

func (da *DefaultAdaptor) On(name string, callback EventCallback) {
	da.Lock()
	da.el[name] = NewEvent(name,callback)
	da.Unlock()
}

func GetDefaultAdaptor() *DefaultAdaptor {
	el := DefaultAdaptor{
		el: make(map[string]*Event, 0),
	}

	return &el
}

type ClientAdaptor interface {
	Call(name string, data string)
	On(name string, callback ClientEventCallback)
	GetEvent(name string) *ClientEvent
}

type ClientDefaultAdaptor struct {
	el map[string]*ClientEvent
	sync.Mutex
}

func GetClientAdapptor() *ClientDefaultAdaptor {
	el := ClientDefaultAdaptor{
		el:make(map[string]*ClientEvent, 0),
	}

	return &el
}

func (da *ClientDefaultAdaptor) Call(name string, data string) {
	e := da.GetEvent(name)
	if e == nil {
		return
	}

	e.callback(data)
}

func (da *ClientDefaultAdaptor) GetEvent(name string) *ClientEvent {
	da.Lock()
	event, ok := da.el[name]
	da.Unlock()

	if !ok {
		return nil
	}

	return event
}

func (da *ClientDefaultAdaptor) On(name string, callback ClientEventCallback) {
	da.Lock()
	da.el[name] = NewClientEvent(name,callback)
	da.Unlock()
}