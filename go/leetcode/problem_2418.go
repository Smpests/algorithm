package leetcode

import "sort"


func sortPeople(names []string, heights []int) []string {
	length := len(names)
	heightsWithIndex := make([][]int, length)
	for i := 0; i < length; i++ {
		heightsWithIndex[i] = make([]int, 2)
		heightsWithIndex[i] = []int{heights[i], i}
	}
	sort.Slice(heightsWithIndex, func(i, j int) bool {
		return heightsWithIndex[j][0] < heightsWithIndex[i][0]
	})
	answer := make([]string, length)
	for i := 0; i < length; i++ {
		answer[i] = names[heightsWithIndex[i][1]]
	}
	return answer
}