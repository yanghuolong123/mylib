package help

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(str string) string {
	s := sha1.New()
	io.WriteString(s, str)

	return fmt.Sprintf("%x", s.Sum(nil))
}
