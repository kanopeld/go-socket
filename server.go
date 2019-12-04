package socket

import (
	"net"
	"sync"
)

type Server interface {
	Start()
	Stop()
	Handler
}

type server struct {
	*baseHandler
	ln net.Listener
	sync.Mutex
	closeChan chan struct{}
}

func NewServer(port string) (Server, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	s := &server{
		baseHandler: newHandler(newDefaultBroadcast(), getCaller("SClient")),
		ln:          ln,
	}
	return s, nil
}

func (s *server) loop() {
	defer func() {
		_ = s.ln.Close()
	}()
	for {
		select {
		case <-s.closeChan:
			return
		default:
			conn, err := s.ln.Accept()
			if err != nil {
				continue
			}
			c, err := newClient(conn, s.baseHandler)
			if err != nil || c == nil {
				continue
			}
			go c.loop()
		}
	}
}

func (s *server) Start() {
	s.Lock()
	defer s.Unlock()
	s.loop()
}

func (s *server) Stop() {
	s.closeChan <- struct{}{}
}
