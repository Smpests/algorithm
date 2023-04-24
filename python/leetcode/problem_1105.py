# 题目链接：https://leetcode.cn/problems/filling-bookcase-shelves/
from typing import List


def minHeightShelves(books: List[List[int]], shelfWidth: int) -> int:
    bookCount = len(books)
    states = [0]
    for book in books:
        states.append(states[-1] + book[1])

    for i in range(bookCount):
        floorMaxHeight, leftWidth = 0, shelfWidth
        for j in range(i, -1, -1):
            leftWidth -= books[j][0]
            if leftWidth < 0:
                break
            floorMaxHeight = max(floorMaxHeight, books[j][1])
            states[i + 1] = min(states[i + 1], states[j] + floorMaxHeight)
    return states[bookCount]
