package socket

const (
	//PackTypeConnect serves for connection messages
	PackTypeConnect PackageType = 0x0
	//PackTypeDisconnect serves for disconnection messages
	PackTypeDisconnect PackageType = 0x1
	//PackTypeEvent serves for event messages
	PackTypeEvent PackageType = 0x2
	//PackTypeConnectAccept serves for connection accepted messages
	PackTypeConnectAccept PackageType = 0x3
)

//PackageType the special type for marking packages
type PackageType byte

//Byte serializes a PackageType into byte
func (pt PackageType) Byte() byte {
	return byte(pt)
}

// Message stores information about a single package
type Package struct {
	PT PackageType
	//Payload usually includes a message
	Payload []byte
}

//MarshalBinary serializes a package into bytes
func (p Package) MarshalBinary() []byte {
	b := make([]byte, 1)
	b[0] = p.PT.Byte()
	if len(p.Payload) > 0 {
		b = append(b, p.Payload...)
	}
	b = append(b, '\n')
	return b
}
