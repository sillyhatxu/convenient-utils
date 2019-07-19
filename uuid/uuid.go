package uuid

import (
	"github.com/oklog/ulid"
	"github.com/satori/go.uuid"
	"github.com/sillyhatxu/sillyhat-cloud-utils/encryption/hash"
	"math/rand"
	"time"
)

func UUIDV4() string {
	u := uuid.Must(uuid.NewV4())
	return u.String()
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
	shortId, err := hash.HashValue32(UUID())
	if err != nil {
		return UUID()
	}
	return randomSrc() + shortId
}
