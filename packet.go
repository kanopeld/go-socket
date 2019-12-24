package socket

// PackageType the special type for marking packages
type PackageType byte

const (
	// PackTypeConnect serves for connection messages
	PackTypeConnect PackageType = iota
	// PackTypeDisconnect serves for disconnection messages
	PackTypeDisconnect
	// PackTypeEvent serves for event messages
	PackTypeEvent
	// PackTypeConnectAccept serves for connection accepted messages
	PackTypeConnectAccept
)

// Byte serializes a PackageType into byte
func (pt PackageType) Byte() byte {
	return byte(pt)
}

// package stores information about a single package
type sockPackage struct {
	// PT is the type of a package
	PT PackageType
	// Payload usually includes a message
	Payload []byte
}

// MarshalBinary serializes a package into bytes
func (p sockPackage) MarshalBinary() []byte {
	b := make([]byte, 1)
	b[0] = p.PT.Byte()
	if len(p.Payload) > 0 {
		b = append(b, p.Payload...)
	}
	b = append(b, '\n')
	return b
}

// decodePackage creates a package from a given message
func decodePackage(msg []byte) sockPackage {
	if msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	p := sockPackage{
		PT:      PackageType(msg[0]),
		Payload: msg[1:],
	}
	return p
}
