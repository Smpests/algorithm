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

func QuickSort[T int|float32|float64|string](data []T, left, right int) {
	if left >= right {
		return
	}
	// 选择中间要快很多，不清楚为什么，都是随机选的数
	pivot := data[(left + right) / 2]
	i, j := left, right
	for i <= j {
		/* 
		// 错误写法
		for i <= j && data[i] <= pivot {
			i++
		}
		for i <= j && data[j] >= pivot {
			j--
		}
		// 此时已经不满足上面两个循环条件，需要交换
		if i < j {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		}
		解释为什么不能写成data[i] <= pivot或data[j] >= pivot，
		也就是为什么快排是Unstable的算法：
		1.如果写成data[i] <= pivot，就是说小于等于基准值的数不用移动，
		此时如果基准值恰好选到数组中的最大值，i在for循环中就会越界，
		因此要加上i <= j的条件；
		2.if i < j，不能是i <= j的原因，因为i = j一定会在上面的循环条件中被覆盖；
		3.当pivot是数组最大值时（或者在某个partition中选到最大值），
		QuickSort(data, left, j)就依然是原输入，不断递归导致栈溢出。

		正确做法就是data[i] < pivot和data[j] > pivot，等于情况也做交换：
		1.此时就不需要i <= j的循环条件，因为即使pivot选到最大值，也会退出循环；
		2.而if判断要写成 i <= j，因为i=j的情况下不满足循环条件，
		则说明data[i]=data[j]=pivot，此时要退出外层循环，需要i++和j--；
		3.如果data[i]=data[j]=pivot，递归时会漏掉这个i=j这个位置的数(i++和j--)，
		但是不影响，因为它已经在正确的位置上；
		4.代码如下。
		*/
		for data[i] < pivot {
			i++
		}
		for data[j] > pivot {
			j--
		}
		if i <= j {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		}
	}
	QuickSort(data, left, j)
	QuickSort(data, i, right)
}