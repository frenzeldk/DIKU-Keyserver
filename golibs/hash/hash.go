package hash

import (
	"crypto/sha512"
	"io"
	//	"strconv"
	//	"time"
)

func GetHash(str1, str2 string) (hash []byte) {
	h512 := sha512.New()
	//	time := time.Now().Unix()        // We use Unix time format
	//	t := strconv.FormatInt(time, 10) // We now change the Int into a String
	secret := "this is our muuch secret string"
	io.WriteString(h512, str1+str2+secret)
	return h512.Sum(nil)
}
