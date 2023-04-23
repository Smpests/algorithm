from typing import List

def mergeSort(data: List[int]) -> List[int]:
    if len(data) <= 1:
        return data
    mid: int = int(len(data) / 2)
    left: List[int] = mergeSort(data[:mid])
    right: List[int] = mergeSort(data[mid:])
    return merge(left, right)

def merge(left: List[int], right: List[int]) -> List[int]:
    result: List[int] = []
    i, j = 0, 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1

    result.extend(left[i:])
    result.extend(right[j:])
    return result

def quickSort(data: List[int], left: int, right: int):
    if left >= right:  # backtrace condition
        return
    pivot = data[int((left + right) / 2)]  # choose pivot
    i, j = left, right
    while i <= j:
        while data[i] < pivot:
            i += 1
        while data[j] > pivot:
            j -= 1
        if i <= j:
            data[i], data[j] = data[j], data[i]
            i += 1
            j -= 1
    quickSort(data, left, j)
    quickSort(data, i, right)
