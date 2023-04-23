import unittest
from python.common.sort import mergeSort, quickSort

class SortUnitTest(unittest.TestCase):

    def test_mergeSort(self):
        arr = [1, 5, 6, 5, 7, 9, 8, 2, 4, 3, 0]
        excepted = arr[:]
        excepted.sort()
        output = mergeSort(arr)
        self.assertEqual(output, excepted)

    def test_quickSort(self):
        arr = [1, 5, 6, 5, 7, 9, 8, 2, 4, 3, 0]
        excepted = arr[:]
        excepted.sort()
        quickSort(arr, 0, len(arr) - 1)
        self.assertEqual(arr, excepted)


if __name__ == '__main__':
    unittest.main()
