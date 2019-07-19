package hashset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSet(t *testing.T) {
	testHashSet := New()
	testHashSet.Add("1")
	testHashSet.Add("2")
	testHashSet.Add("3")
	testHashSet.Add("A")
	testHashSet.Add("B")
	testHashSet.Add("C")
	testHashSet.Add("2")
	testHashSet.Add("C")
	assert.NotNil(t, testHashSet)
	assert.EqualValues(t, testHashSet.Size(), 6)
	checkArray := [6]string{"1", "2", "3", "A", "B", "C"}
	for _, src := range testHashSet.ToArray() {
		check := false
		for _, checkSrc := range checkArray {
			if src == checkSrc {
				check = true
			}
		}
		assert.EqualValues(t, check, true)

	}
}
