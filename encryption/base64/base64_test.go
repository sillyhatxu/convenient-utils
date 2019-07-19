package base64

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	msg    = "Hello, 世界"
	encode = "SGVsbG8sIOS4lueVjA=="
)

func TestEncodeToString(t *testing.T) {
	test := EncodeToString([]byte(msg))
	assert.EqualValues(t, test, encode)
}

func TestDecodeString(t *testing.T) {
	test, err := DecodeString(encode)
	assert.Nil(t, err)
	assert.EqualValues(t, test, msg)
}
