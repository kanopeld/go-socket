package socket

import "net"

type sender interface {
	send(p *sockPackage) error
}

func getSender(c net.Conn) sender {
	return &send{c: c}
}

type send struct {
	c net.Conn
}

func (s *send) send(p *sockPackage) error {
	_, err := s.c.Write(p.MarshalBinary())
	return err
}
