package remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Given nums = [3,2,2,3], val = 3,
// Your function should return length = 2, with the first two elements of nums being 2.
// It doesn't matter what you leave beyond the returned length.
func TestExample1(t *testing.T) {
	result := RemoveElement([]int{3, 2, 2, 3}, 3)

	assert.Equal(t, 2, result)
}

// Given nums = [0,1,2,2,3,0,4,2], val = 2,
// Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.
// Note that the order of those five elements can be arbitrary.
// It doesn't matter what values are set beyond the returned length.
func TestExample2(t *testing.T) {
	result := RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)

	assert.Equal(t, 5, result)
}
