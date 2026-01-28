package utils

import (
	"crypto/md5"
	"encoding/base64"
)

func GetHashPassword(password string) string {
	hashed5 := md5.Sum([]byte(password))
	hashed5Str := base64.StdEncoding.EncodeToString(hashed5[:])
	return hashed5Str
}
