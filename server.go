package socket

import (
	"net"
	"sync"
)

// Server is an object managing your connection
type server struct {
	*baseHandler
	ln net.Listener
	sync.Mutex
	closeChan chan struct{}
}

// NewServer starts a tcp listener on a specified port and
// returns an initialized Server struct
func NewServer(port string) (*server, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	s := &server{
		baseHandler: newHandler(newDefaultBroadcast()),
		ln:          ln,
		closeChan:   make(chan struct{}),
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

// Start starts a server
func (s *server) Start() {
	s.Lock()
	defer s.Unlock()
	s.loop()
}

// Stop stops a server
func (s *server) Stop() {
	s.closeChan <- struct{}{}
}
