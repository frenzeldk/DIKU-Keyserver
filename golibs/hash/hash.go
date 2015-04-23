package hash

import (
	"crypto/sha512"
	"io"
)

func GetHash(str string) (hash []byte) {
	h512 := sha512.New()
	secret := "this is our muuch secret string"
	io.WriteString(h512, str + secret)
	return h512.Sum(nil)
}
