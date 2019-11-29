package socket

// PackageType is an enum for packet types
type PackageType byte

const (
	_PACKET_TYPE_CONNECT PackageType = iota
	_PACKET_TYPE_DISCONNECT
	_PACKET_TYPE_EVENT
	_PACKET_TYPE_CONNECT_ACCEPT
)

func (pt PackageType) byte() byte {
	return byte(pt)
}

// Package stores some payload and its type
type Package struct {
	PT      PackageType
	Payload []byte
}

// MarshalBinary packs a Package into a single byte array
func (p Package) MarshalBinary() []byte {
	b := make([]byte, 1)
	b[0] = p.PT.byte()
	if len(p.Payload) > 0 {
		b = append(b, p.Payload...)
	}
	b = append(b, '\n')

	return b
}
