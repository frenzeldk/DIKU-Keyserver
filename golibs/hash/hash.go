package hash

import (
	"crypto/sha512"
	"io"
	"strconv"
	"time"
)

func GetHash(str string) (hash []byte) {
	h512 := sha512.New()
	time := time.Now().Unix()        // We use Unix time format
	t := strconv.FormatInt(time, 10) // We now change the Int into a String
	io.WriteString(h512, t+str)
	return h512.Sum(nil)
}
