package server

import (
	"github.com/kanopeld/go-socket/core"
	"net"
	"sync"
)

type Server struct {
	*core.BaseHandler
	ln net.Listener
	sync.Mutex
	closeChan chan struct{}
}

func NewServer(port string) (*Server, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	s := &Server{
		BaseHandler: core.NewHandler(core.NewDefaultBroadcast(), NewCaller),
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
			c, err := newClient(conn, s.BaseHandler)
			if err != nil || c == nil {
				continue
			}
			go c.Loop()
		}
	}
}

func (s *Server) Start() {
	s.loop()
}

func (s *Server) Stop() {
	s.closeChan <- struct{}{}
}
