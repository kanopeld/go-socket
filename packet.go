package socket

const (
	_PACKET_TYPE_CONNECT PackageType = 0x00
	_PACKET_TYPE_DISCONNECT PackageType = 0x01
	_PACKET_TYPE_EVENT PackageType = 0x02
)

type PackageType byte

func (pt PackageType) byte() byte {
	switch pt {
	case _PACKET_TYPE_CONNECT:
		return 0x00
	case _PACKET_TYPE_DISCONNECT:
		return 0x01
	case _PACKET_TYPE_EVENT:
		return 0x02
	}

	return 0x03
}

type Package struct {
	PT PackageType `json:"t"`
	Payload []byte `json:"p"`
}

func (p *Package) MarshalBinary() []byte {
	b := make([]byte, 1)
	b[0] = p.PT.byte()
	if len(p.Payload) > 0 {
		b = append(b, p.Payload...)
	}
	b = append(b, '\n')

	return b
}

func NewPacket(pt PackageType) Package {
	return Package{PT:pt}
}
