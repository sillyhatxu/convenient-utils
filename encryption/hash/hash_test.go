package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashError(t *testing.T) {
	src := "error"
	hashSrc32, err := HashValue32(src)
	assert.Nil(t, err)
	assert.EqualValues(t, hashSrc32, "563185489")
}
func TestHash(t *testing.T) {
	src := "2019"
	hashSrc32, err := HashValue32(src)
	assert.Nil(t, err)
	hashSrc64, err := HashValue64(src)
	assert.Nil(t, err)
	assert.EqualValues(t, hashSrc32, "3308014985")
	assert.EqualValues(t, hashSrc64, "1745884875074361097")
}
