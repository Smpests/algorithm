package common


func MergeSort(data []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	MergeSort(data, left, mid)
	MergeSort(data, mid + 1, right)
	// 回溯时合并，要合并的区间是[left, right],两两合并
	// 两个子数组分别是：[left, mid]、[mid + 1, right]
	mergedSlice := make([]int, 0, right - left + 1)
	// 两个子数组起始位置在原数组中的索引
	firstCursor, secondCursor := left, mid + 1
	for i := 0; i < right; i++ {
		// 任意一边的子数组比较到边界，剩下的已经有序，则退出
		if firstCursor > mid {
			mergedSlice = append(mergedSlice, data[secondCursor])
			secondCursor += 1
		} else if secondCursor > right {
			mergedSlice = append(mergedSlice, data[firstCursor])
			firstCursor += 1
		} else if data[firstCursor] <= data[secondCursor] {
			mergedSlice = append(mergedSlice, data[firstCursor])
			firstCursor += 1
		} else {
			mergedSlice = append(mergedSlice, data[secondCursor])
			secondCursor += 1
		}
	}
	for i := left; i < right + 1; i++ {
		data[i] = mergedSlice[i - left]
	}
}