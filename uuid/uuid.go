package uuid

import (
	"github.com/oklog/ulid"
	"github.com/satori/go.uuid"
	"github.com/sillyhatxu/convenient-utils/encryption/hash"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func V4() string {
	return uuid.NewV4().String()
}

func V4Upper32() string {
	return strings.ToUpper(strings.ReplaceAll(uuid.NewV4().String(), "-", ""))
}

func UUID() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomSrc() string {
	return string(letters[rand.Intn(len(letters))])
}

func ShortId() string {
	shortId, err := hash.Hash32(UUID())
	if err != nil {
		return UUID()
	}
	return randomSrc() + strconv.FormatUint(uint64(shortId), 10)
}
