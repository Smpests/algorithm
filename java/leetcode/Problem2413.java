package leetcode;

/**
 * 题目链接：https://leetcode.cn/problems/smallest-even-multiple/
 */
class Problem2413 {
    public int smallestEvenMultiple(int n) {
        if (n % 2 == 0) {
            return n;
        }
        return 2 * n;
    }
    public static void main(String[] args) {
        Problem2413 problem = new Problem2413();
        int result = problem.smallestEvenMultiple(5);
        System.out.println(result);
    }
}
