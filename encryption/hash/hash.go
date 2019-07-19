package hash

import (
	"hash/fnv"
	"strconv"
)

func HashValue32(s string) (string, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(uint64(h.Sum32()), 10), nil
}

func HashValue64(s string) (string, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(h.Sum64(), 10), nil
}
