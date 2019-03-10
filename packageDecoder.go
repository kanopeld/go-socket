package socket

const (
	CONNECTION_NAME = "connection"
	DISCONNECTION_NAME = "disconnection"
)

func DecodePackage(msg []byte) (Package, error) {
	p := Package{
		PT:PackageType(msg[0]),
		Payload:msg[1:],
	}

	return p, nil
}