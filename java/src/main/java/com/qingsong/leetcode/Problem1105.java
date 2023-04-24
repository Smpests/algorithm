package com.qingsong.leetcode;

import java.lang.Math;


class Problem1105 {
    public int minHeightShelves(int[][] books, int shelfWidth) {
        int[] states = new int[books.length + 1];
        for (int i = 0; i < books.length; ++i) {
            states[i + 1] = states[i] + books[i][1];
        }
        
        for (int i = 0; i < books.length; ++i) {
            int floorMaxHeight = 0;
            int leftWidth = shelfWidth;
            for (int j = i; j > -1; --j) {
                leftWidth -= books[j][0];
                if (leftWidth < 0) {
                    break;
                }
                floorMaxHeight = Math.max(floorMaxHeight, books[j][1]);
                states[i + 1] = Math.min(states[i + 1], states[j] + floorMaxHeight);
            }
        }
        return states[books.length];
    }
}