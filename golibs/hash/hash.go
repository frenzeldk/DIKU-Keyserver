package hash

import (
	"crypto/sha224"
	"io"
)

func GetHash(str string) (hash []byte) {
	h224 := sha224.New()
	secret := "this is our muuch secret string"
	io.WriteString(h224, str + secret)
	return h224.Sum(nil)
}
