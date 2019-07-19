package base64

import (
	"encoding/base64"
)

func DecodeString(s string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func EncodeToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
