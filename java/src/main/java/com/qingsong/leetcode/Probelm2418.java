package com.qingsong.leetcode;

import java.util.Arrays;


/**
 * 题目链接：https://leetcode.cn/problems/sort-the-people/
 */
class Problem2418 {
    public String[] sortPeople(String[] names, int[] heights) {
        int length = names.length;
        int[][] heightsWithIdx = new int[length][2];  // {{height, index}, ...}
        for (int i = 0; i < length; ++i) {
            heightsWithIdx[i][0] = heights[i];
            heightsWithIdx[i][1] = i;
        }
        Arrays.sort(heightsWithIdx, (x, y) -> y[0] - x[0]);  // secondParameter - firstParameter for reverse order
        String[] answer = new String[length];
        for (int i = 0; i < length; ++i) {
            answer[i] = names[heightsWithIdx[i][1]];
        }
        return answer;
    }

    public static void main(String[] args) {
        Problem2418 problem = new Problem2418();
        String[] names = {"Mary","John","Emma"};
        int[] heights = {180,165,170};
        problem.sortPeople(names, heights);
    }
}