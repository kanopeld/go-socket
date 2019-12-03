package socket

// DecodePackage creates a package from a given message
func DecodePackage(msg []byte) (Package, error) {
	if msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	p := Package{
		PT:      PackageType(msg[0]),
		Payload: msg[1:],
	}
	return p, nil
}
