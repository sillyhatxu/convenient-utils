package md5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	src := "2019"
	md5 := Encrypt(src)
	assert.EqualValues(t, len(md5), 32)
	assert.EqualValues(t, md5, "32303139d41d8cd98f00b204e9800998")
	md5Upper := EncryptUpper(src)
	assert.EqualValues(t, len(md5Upper), 32)
	assert.EqualValues(t, md5Upper, "32303139D41D8CD98F00B204E9800998")
	md5LongUpper := EncryptUpper("xushikuanissillyhatthisislongsrc")
	assert.EqualValues(t, len(md5LongUpper), 32)
	assert.EqualValues(t, md5LongUpper, "78757368696B75616E697373696C6C79")
}

func TestMD5Temp(t *testing.T) {
	src := "123456"
	md5 := EncryptUpper(src)
	fmt.Println(md5)
}
