package md5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"unicode"
)

func Encrypt(message string) string {
	hash := md5.New()
	hashInBytes := hash.Sum([]byte(message))[:16]
	return hex.EncodeToString(hashInBytes)
}

func EncryptUpper(message string) string {
	hash := md5.New()
	hashInBytes := hash.Sum([]byte(message))[:16]
	return strings.ToUpperSpecial(unicode.TurkishCase, hex.EncodeToString(hashInBytes))
}
