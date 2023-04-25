package leetcode

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestSmallestEvenMultiple(t *testing.T) {
	assert.Equal(t, 10, smallestEvenMultiple(5))
}

func TestMinHeightShelves(t *testing.T) {
	expected := 6
	input := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
	output := minHeightShelves(input, 4)
	assert.Equal(t, expected, output)
}

func TestSortPeople(t *testing.T) {
	names := []string{"Mary","John","Emma"}
	heights := []int{180,165,170}
	expected := []string{"Mary","Emma","John"}
	output := sortPeople(names, heights)
	assert.Equal(t, expected, output)
}