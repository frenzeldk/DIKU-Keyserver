package hash

import (
	"crypto/sha512"
	"io"
	"strconv"
	"time"
)

func GetHash(kuid string) (hash []byte) {
	h512 := sha512.New()
	tid := time.Now().Unix()        // We use Unix time format
	t := strconv.FormatInt(tid, 10) // We now change the Int into a String
	io.WriteString(h512, t+kuid)
	return h512.Sum(nil)
}
