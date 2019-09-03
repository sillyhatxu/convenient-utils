package uuid

import (
	"fmt"
	"github.com/sillyhatxu/convenient-utils/hashset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortId(t *testing.T) {
	idSet := hashset.New()
	for i := 0; i < 100000; i++ {
		id := ShortId()
		fmt.Println(id)
		idSet.Add(id)
	}
	assert.EqualValues(t, idSet.Size(), 100000)
}

func TestGetShortId(t *testing.T) {
	id := ShortId()
	fmt.Println(id)
}

func TestUUID(t *testing.T) {
	idSet := hashset.New()
	for i := 0; i < 100000; i++ {
		id := V4Upper32()
		fmt.Println(id)
		idSet.Add(id)
	}
	assert.EqualValues(t, idSet.Size(), 100000)
}
