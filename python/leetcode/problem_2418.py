from typing import List


# 题目链接：https://leetcode.cn/problems/sort-the-people/
def sortPeople(names: List[str], heights: List[int]) -> List[str]:
    return [name for name, _ in sorted(zip(names, heights), key=lambda x: -x[1])]
