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

func TestDinnerPlates(t *testing.T) {
	dinnerPlates := DinnerPlates{capacity: 2}
	dinnerPlates.Push(1)
	dinnerPlates.Push(2)
	dinnerPlates.Push(3)
	dinnerPlates.Push(4)
	dinnerPlates.Push(5)
	assert.Equal(t, 2, dinnerPlates.PopAtStack(0))
	dinnerPlates.Push(20)
	dinnerPlates.Push(21)
	assert.Equal(t, 20, dinnerPlates.PopAtStack(0))
	assert.Equal(t, 21, dinnerPlates.PopAtStack(2))
	assert.Equal(t, 5, dinnerPlates.Pop())
	assert.Equal(t, 4, dinnerPlates.Pop())
	assert.Equal(t, 3, dinnerPlates.Pop())
	assert.Equal(t, 1, dinnerPlates.Pop())
	assert.Equal(t, -1, dinnerPlates.Pop())
}