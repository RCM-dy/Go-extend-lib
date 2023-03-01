package Go_extend_lib

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Bytes(datas []byte) string {
	s := sha1.New()
	s.Write(datas)
	return hex.EncodeToString(s.Sum(nil))
}
