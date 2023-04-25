import unittest
from python.leetcode import *

class LeetcodeTestCase(unittest.TestCase):

    def testSmallestEvenMultiple(self):
        self.assertEqual(10, smallestEvenMultiple(5))

    def testMinHeightShelves(self):
        _input = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]]
        self.assertEqual(6, minHeightShelves(_input, 4))

    def testSortPeople(self):
        names = ["Mary","John","Emma"];
        heights = [180,165,170]
        expected = ["Mary","Emma","John"]
        self.assertListEqual(expected, sortPeople(names, heights))


if __name__ == '__main__':
    unittest.main()
