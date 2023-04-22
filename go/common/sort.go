package common

func MergeSort[T int|float32|float64|string](data []T) []T {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return merge(left, right)
}

func merge[T int|float32|float64|string](left, right []T) []T {
	leftLen, rightLen := len(left), len(right)
	result := make([]T, 0, leftLen + rightLen)  // new slice for sorted elements
	i, j := 0, 0  // index of left slice: i, index of right slice: j
	for i < leftLen && j < rightLen {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// append the rest elements in left or right
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}