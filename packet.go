package socket

const (
	PackTypeConnect       PackageType = 0x0
	PackTypeDisconnect    PackageType = 0x1
	PackTypeEvent         PackageType = 0x2
	PackTypeConnectAccept PackageType = 0x3
)

type PackageType byte

func (pt PackageType) byte() byte {
	return byte(pt)
}

type Package struct {
	PT      PackageType
	Payload []byte
}

func (p Package) MarshalBinary() []byte {
	b := make([]byte, 1)
	b[0] = p.PT.byte()
	if len(p.Payload) > 0 {
		b = append(b, p.Payload...)
	}
	b = append(b, '\n')
	return b
}
