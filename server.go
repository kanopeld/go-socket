package socket

import (
	"net"
	"sync"
)

// Server is an object managing your connection
type Server struct {
	*baseHandler
	ln net.Listener
	sync.Mutex
	closeChan chan struct{}
}

// NewServer starts a tcp listener on a specified port and
// returns an initialized Server struct
func NewServer(port string) (*Server, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	s := &Server{
		baseHandler: newHandler(newDefaultBroadcast()),
		ln:          ln,
		closeChan:   make(chan struct{}),
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

// Start starts a server
func (s *Server) Start() {
	s.Lock()
	defer s.Unlock()
	s.loop()
}

// Stop stops a server
func (s *Server) Stop() {
	s.closeChan <- struct{}{}
}
