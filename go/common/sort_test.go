package common

import (
	"testing"
	"math/rand"
	"time"
	"sort"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	input := []int{10, 2, 3, 1, 5, 4, 6, 9, 8, 7}
	excepted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	input = MergeSort(input)
	assert.Equal(t, excepted, input, "Wrong order")
}

func generateRandomIntSlice(n, maxValue int) []int {
	testData := make([]int, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:= 0; i < n; i++ {
		testData[i] = r.Intn(maxValue)
	}
	return testData
}

func BenchmarkMergeSort(b *testing.B) {
	// 72	  16477418 ns/op	14057769 B/op	   99999 allocs/op
	// 空间开辟消耗大，比sort.Ints()性能低很多
	testData := generateRandomIntSlice(100000, 50000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(testData)
	}
}

func BenchmarkSortInts(b *testing.B) {
	testData := generateRandomIntSlice(100000, 50000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(testData)
	}
}
