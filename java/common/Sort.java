package common;

import java.util.Arrays;

public class Sort {

    public static int[] mergeSort(int[] data) {
        int dataLen = data.length;
        if (data.length <= 1) {
            return data;
        }
        int mid = data.length / 2;
        int[] left = mergeSort(Arrays.copyOfRange(data, 0, mid));
        int[] right = mergeSort(Arrays.copyOfRange(data, mid, dataLen));
        return merge(left, right);
    }

    public static int[] merge(int[] left, int[] right) {
        int[] result = new int[left.length + right.length];
        int i = 0;
        int j = 0;
        int k = 0;
        while (i < left.length && j < right.length) {
            if (left[i] <= right[j]) {
                result[k++] = left[i++];
            } else {
                result[k++] = right[j++];
            }
        }
        while (i < left.length) {
            result[k++] = left[i++];
        }
        while (j < right.length) {
            result[k++] = right[j++];
        }
        return result;
    }

    public static void quickSort(int[] data, int left, int right) {
        if (left >= right) {
            return;
        }
        int pivot = data[(left + right) / 2];
        int i = left;
        int j = right;
        while (i <= j) {
            while (data[i] < pivot) {
                ++i;
            }
            while (data[j] > pivot) {
                --j;
            }
            if (i <= j) {
                int temp = data[i];
                data[i] = data[j];
                data[j] = temp;
                ++i;
                --j;
            }
        }
        quickSort(data, left, j);
        quickSort(data, i, right);
    }

    public static void main(String[] args) {
        int[] input = new int[] {10, 2, 3, 1, 5, 4, 6, 9, 8, 7};
        int[] excepted = new int[] {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

        int[] outputMergeSort = Sort.mergeSort(input);
        for (int i = 0; i < input.length; ++i) {
            if (outputMergeSort[i] != excepted[i]) {
                System.out.println("mergeSort(): wrong order");
                return;
            }
        }
        System.out.println("mergeSort() pass");
        quickSort(input, 0, input.length - 1);
        for (int i = 0; i < input.length; ++i) {
            if (input[i] != excepted[i]) {
                System.out.println("quickSort(): wrong order");
                return;
            }
        }
        System.out.println("quickSort() pass");
    }
}