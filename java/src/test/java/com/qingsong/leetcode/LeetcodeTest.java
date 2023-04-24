package com.qingsong.leetcode;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;



/**
 * Unit test for simple Leetcode.
 */
public class LeetcodeTest {
    /**
     * Rigorous Test :-)
     */
    @Test
    public void testSmallestEvenMultiple() {
        Problem2413 problem = new Problem2413();
        Assertions.assertEquals(10, problem.smallestEvenMultiple(5));
    }

    @Test
    public void testMinHeightShelves() {
        Problem1105 problem = new Problem1105();
        int[][] input = {{1, 1,}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}};
        Assertions.assertEquals(6, problem.minHeightShelves(input, 4));
    }
}
