package hash

import (
	"crypto/sha256"
	"io"
)

func GetHash(str string) (hash []byte) {
	h256 := sha256.New()
	secret := "this is our muuch secret string"
	io.WriteString(h256, str + secret)
	return h256.Sum(nil)
}
