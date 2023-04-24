package leetcode

// 题目链接：https://leetcode.cn/problems/smallest-even-multiple/
func smallestEvenMultiple(n int) int {
    if n % 2 == 0 {  // 是偶数
        return n
    }
    // 比写else快很多
    return n + n
}