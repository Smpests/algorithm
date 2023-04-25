package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSearch	(t *testing.T) {
	var segmentTree SegmentTree
	arr := []int{1, 2, 3, 4, 5}
	segmentTree.Build(arr)
	assert.Equal(t, 7, segmentTree.Query(2, 3))
}

func TestUpdate(t *testing.T) {
	var segmentTree SegmentTree
	arr := []int{1, 2, 3, 4, 5}
	segmentTree.Build(arr)
	segmentTree.Update(2, 6)
	assert.Equal(t, 10, segmentTree.Query(2, 3))
}

func TestBigSliceBuild(t *testing.T) {
	var segmentTree SegmentTree
	randomSlice := generateRandomIntSlice(100000, 1000)
	segmentTree.Build(randomSlice)
	// average took 0.03s
}