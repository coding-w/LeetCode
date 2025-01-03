package sort;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.PriorityQueue;
import java.util.Random;


// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 冒泡排序
        System.out.println("bubbleSort = " + Arrays.toString(bubbleSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 选择排序
        System.out.println("selectionSort = " + Arrays.toString(selectionSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 插入排序
        System.out.println("insertionSort = " + Arrays.toString(insertionSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 希尔排序
        System.out.println("shellSort = " + Arrays.toString(shellSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 归并排序
        System.out.println("mergeSort = " + Arrays.toString(mergeSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 快速排序
        System.out.println("quickSort = " + Arrays.toString(quickSort(new int[]{3, 5, 4, 6, 1, 2}, 0, 5)));

        // 堆排序
        System.out.println("heapSort = " + Arrays.toString(heapSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 计数排序
        System.out.println("countingSort = " + Arrays.toString(countingSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 桶排序
        System.out.println("bucketSort = " + Arrays.toString(bucketSort(new int[]{3, 5, 4, 6, 2, 1})));

        // 基数排序
        System.out.println("radixSort = " + Arrays.toString(radixSort(new int[]{113, 112, 210, 6})));

        // timSort
        System.out.println("timSort = " + Arrays.toString(timSort(new int[]{3, 5, 4, 6, 2, 1})));
        System.out.println("bogoSort = " + Arrays.toString(bogoSort(new int[]{3, 5, 4, 6, 2, 1})));
    }




    // bubbleSort 冒泡排序
    public static int[] bubbleSort(int[] nums) {
        for (int i = 0; i < nums.length - 1; i++) {
            for (int j = 1; j < nums.length - i; j++) {
                if (nums[j] < nums[j - 1]) {
                    int temp = nums[j];
                    nums[j] = nums[j - 1];
                    nums[j - 1] = temp;
                }
            }
        }
        return nums;
    }

    // selectionSort 选择排序
    public static int[] selectionSort(int[] nums) {
        for (int i = 0; i < nums.length - 1; i++) {
            int minIndex = i;
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[j] < nums[minIndex]) {
                    minIndex = j;
                }
            }
            if (minIndex != i) {
                int temp = nums[minIndex];
                nums[minIndex] = nums[i];
                nums[i] = temp;
            }
        }
        return nums;
    }

    // insertionSort 插入排序
    public static int[] insertionSort(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            int temp = nums[i];
            int j = i - 1;
            while (j >= 0 && nums[j] > temp) {
                nums[j + 1] = nums[j];
                j--;
            }
            nums[j + 1] = temp;
        }
        return nums;
    }

    // shellSort 希尔排序
    public static int[] shellSort(int[] nums) {
        int gap = nums.length / 2;
        while (gap > 0) {
            for (int i = gap; i < nums.length; i++) {
                int temp = nums[i];
                int j = i - gap;
                while (j >= 0 && nums[j] > temp) {
                    nums[j + gap] = nums[j];
                    j -= gap;
                }
                nums[j + gap] = temp;
            }
            gap /= 2;
        }
        return nums;
    }

    // mergeSort 归并排序
    public static int[] mergeSort(int[] nums) {
        if (nums.length < 2) {
            return nums;
        }
        int mid = nums.length / 2;
        int[] left = new int[mid];
        int[] right = new int[nums.length - mid];
        System.arraycopy(nums, 0, left, 0, left.length);
        System.arraycopy(nums, mid, right, 0, right.length);
        left = mergeSort(left);
        right = mergeSort(right);
        return merge(left, right);
    }

    // 合并
    public static int[] merge(int[] left, int[] right) {
        int[] merged = new int[left.length + right.length];
        int i = 0, j = 0, k = 0;
        while (i < left.length && j < right.length) {
            if (left[i] <= right[j]) {
                merged[k++] = left[i++];
            } else {
                merged[k++] = right[j++];
            }
        }
        while (i < left.length) {
            merged[k++] = left[i++];
        }
        while (j < right.length) {
            merged[k++] = right[j++];
        }
        return merged;
    }

    // quickSort 快速排序
    public static int[] quickSort(int[] nums, int left, int right) {
        if (left < right) {
            int index = partition(nums, left, right);
            quickSort(nums, left, index - 1);
            quickSort(nums, index + 1, right);
        }
        return nums;
    }

    private static int partition(int[] nums, int left, int right) {
        int pivot = nums[left];
        int i = left;
        for (int j = left + 1; j <= right; j++) {
            if (nums[j] <= pivot) {
                i++;
                int temp = nums[j];
                nums[j] = nums[i];
                nums[i] = temp;
            }
        }
        int temp = nums[left];
        nums[left] = nums[i];
        nums[i] = temp;
        return i;
    }

    // heapSort 堆排序
    public static int[] heapSort(int[] nums) {
        PriorityQueue<Integer> heap = new PriorityQueue<>();
        for (int num : nums) {
            heap.add(num);
        }
        // 从堆中逐个取出元素，按顺序填充到数组
        for (int i = 0; i < nums.length; i++) {
            nums[i] = heap.poll(); // 取出堆顶元素并移除
        }
        return nums;
    }

    // countingSort 计数排序
    public static int[] countingSort(int[] nums) {
        if (nums == null || nums.length == 0) {
            return nums;
        }
        int maxNum = nums[0], minNum = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < minNum) {
                minNum = nums[i];
            }
            if (nums[i] > maxNum) {
                maxNum = nums[i];
            }
        }
        int[] count = new int[maxNum - minNum + 1];
        for (int i = 0; i < count.length; i++) {
            count[nums[i] - minNum]++;
        }
        int index = 0;
        for (int i = 0; i < count.length; i++) {
            while (count[i]-- > 0) {
                nums[index] = i + minNum;
                index++;
            }
        }
        return nums;
    }

    // bucketSort 桶排序
    public static int[] bucketSort(int[] nums) {
        if (nums == null || nums.length <= 1) {
            return nums;
        }

        // 找出数组的最大值和最小值
        int maxNum = nums[0], minNum = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < minNum) {
                minNum = nums[i];
            }
            if (nums[i] > maxNum) {
                maxNum = nums[i];
            }
        }

        // 创建桶，桶的数量可以自定义，通常取决于数据的分布
        int bucketCount = (maxNum - minNum) / nums.length + 1;  // 保证桶的数量适合数据分布
        ArrayList<Integer>[] buckets = new ArrayList[bucketCount];

        // 将元素分配到桶中
        for (int num : nums) {
            int index = (num - minNum) * (bucketCount - 1) / (maxNum - minNum);
            if (buckets[index] == null) {
                buckets[index] = new ArrayList<>();
            }
            buckets[index].add(num);
        }

        // 对每个桶内的元素进行排序，并合并所有桶中的元素
        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < bucketCount; i++) {
            if (buckets[i] != null && !buckets[i].isEmpty()) {
                Collections.sort(buckets[i]);
                list.addAll(buckets[i]);
            }
        }

        // 将排序后的结果放回原数组
        for (int i = 0; i < nums.length; i++) {
            nums[i] = list.get(i);
        }

        return nums;
    }

    // radixSort 基数排序
    public static int[] radixSort(int[] nums) {
        if (nums == null || nums.length == 0) {
            return nums;
        }
        int maxNum = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (maxNum < nums[i]) {
                maxNum = nums[i];
            }
        }
        int exp = 1;
        while (maxNum / exp > 0) {
            nums = countingSortByDigit(nums, exp);
            exp *= 10;
        }
        return nums;
    }

    // countingSortByDigit 对数组的某一位进行计数排序
    private static int[] countingSortByDigit(int[] nums, int exp) {
        int[] count = new int[10];
        // 统计每一个数字出现的次数
        for (int num : nums) {
            int index = (num / exp) % 10;
            count[index]++;
        }

        // 计算每个数字的累积计数
        for (int i = 1; i < count.length; i++) {
            count[i] += count[i - 1];
        }
        int[] result = new int[nums.length];
        for (int i = nums.length - 1; i >= 0; i--) {
            int index = (nums[i] / exp) % 10;
            result[count[index] - 1] = nums[i];
            count[index]--;
        }
        return result;
    }

    // timSort
    public static int[] timSort(int[] nums) {
        int n = nums.length;
        final int run = 32;
        for (int i = 0; i < n; i+=run) {
            insertionSortRange(nums, i, Math.min(i+run, n));
        }
        // 合并
        for (int size = run; size < n; size*=2) {
            for (int left = 0; left < n; left+=2*size) {
                int mid = Math.min(n - 1, left + size - 1);
                int right = Math.min((left + 2 * size - 1), (n - 1));
                if (mid < right) {
                    merge(nums, left, mid, right);
                }
            }
        }
        return nums;
    }

    // 合并两个已排序的区间
    private static void merge(int[] arr, int left, int mid, int right) {
        // 如果两个区间已经是有序的，直接返回
        if (arr[mid] <= arr[mid + 1]) return;

        int len1 = mid - left + 1;
        int len2 = right - mid;

        // 临时数组
        int[] leftArray = new int[len1];
        int[] rightArray = new int[len2];

        System.arraycopy(arr, left, leftArray, 0, len1);
        System.arraycopy(arr, mid + 1, rightArray, 0, len2);

        int i = 0, j = 0, k = left;

        // 合并过程
        while (i < len1 && j < len2) {
            if (leftArray[i] <= rightArray[j]) {
                arr[k++] = leftArray[i++];
            } else {
                arr[k++] = rightArray[j++];
            }
        }

        // 如果左半部分剩余元素
        while (i < len1) {
            arr[k++] = leftArray[i++];
        }

        // 如果右半部分剩余元素
        while (j < len2) {
            arr[k++] = rightArray[j++];
        }
    }

    private static void insertionSortRange(int[] nums, int left, int right) {
        for (int i = left + 1; i < right; i++) {
            int temp = nums[i];
            int j = i-1;
            while (j >= left && nums[j] > temp) {
                nums[j+1] = nums[j--];
            }
            nums[j+1] = temp;
        }
    }

    // 检查数组是否已排序
    public static boolean isSorted(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            if (nums[i - 1] > nums[i]) {
                return false;
            }
        }
        return true;
    }

    // 随机打乱数组
    public static void shuffle(int[] nums) {
        Random random = new Random();
        for (int i = 0; i < nums.length; i++) {
            int j = random.nextInt(nums.length); // 随机选择一个索引
            // 交换 arr[i] 和 arr[j] 的元素
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
        }
    }

    // bogoSort 猴子排序
    public static int[] bogoSort(int[] nums) {
        while (!isSorted(nums)) {
            shuffle(nums); // 打乱数组
        }
        return nums;
    }


}
