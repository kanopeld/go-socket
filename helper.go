package socket

import (
	"unicode/utf8"
	"crypto/md5"
	"encoding/hex"
)

func CountStr(str string) int {
	count := 0
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		count++
		str = str[:len(str)-size]
	}

	return count
}

func GetHash(st string) string {
	hasher := md5.New()
	hasher.Write([]byte(st))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}
