package socket

import (
	"net"
	"sync"
)

// Server controls the server side of a connection
type Server struct {
	*baseHandler
	ln net.Listener
	sync.Mutex
	closeChan chan struct{}
}

// NewServer initializes a server
func NewServer(p string) (*Server, error) {
	ln, err := net.Listen("tcp", p)
	if err != nil {
		return nil, err
	}

	s := &Server{
		baseHandler: newBaseHandler(newDefaultBroadcast()),
		ln:          ln,
	}
	return s, nil
}

func (s *Server) loop() {
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

// Start starts the server
func (s *Server) Start() {
	s.loop()
}

// Stop stops the server
func (s *Server) Stop() {
	s.closeChan <- struct{}{}
}
