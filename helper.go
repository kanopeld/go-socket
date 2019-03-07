package socket

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHash(st string) string {
	hasher := md5.New()
	hasher.Write([]byte(st))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}
