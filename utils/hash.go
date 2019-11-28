package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "slKJGt"

// md5加密
func MD5Pwd(val string) string {
	m5 := md5.New()
	m5.Write([]byte(val))
	m5.Write([]byte(string(salt)))
	rt := m5.Sum(nil)
	return hex.EncodeToString(rt)
}
