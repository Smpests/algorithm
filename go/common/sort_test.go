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

func TestQuickSort(t *testing.T) {
	input := []int{10, 2, 3, 1, 5, 4, 6, 9, 8, 7}
	excepted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	QuickSort(input, 0, len(input) - 1)
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
	// 同一个输入数组没必要跑多次，但是把随机数生成放在for循环内跑半天没结果
	for i := 0; i < b.N; i++ {
		MergeSort(testData)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	// pivot选择第一个数时比MergeSort耗时长
	// 因为快排的时间复杂度会随数据在O(nlogn)-O(n²)之间波动
	// pivot选择中间时则快很多，如下：
	// 616	   1980384 ns/op	       0 B/op	       0 allocs/op
	testData := generateRandomIntSlice(100000, 50000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(testData, 0, len(testData) - 1)
	}
}

func BenchmarkSortInts(b *testing.B) {
	// 3866	    313973 ns/op	      24 B/op	       1 allocs/op
	testData := generateRandomIntSlice(100000, 50000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(testData)
	}
}


