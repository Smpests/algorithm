import unittest
from python.leetcode import *

class LeetcodeTestCase(unittest.TestCase):

    def testSmallestEvenMultiple(self):
        self.assertEqual(10, smallestEvenMultiple(5))

    def testMinHeightShelves(self):
        _input = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]]
        self.assertEqual(6, minHeightShelves(_input, 4))


if __name__ == '__main__':
    unittest.main()
