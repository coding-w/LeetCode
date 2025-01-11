package t373;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.PriorityQueue;

class Solution {
    public List<List<Integer>> kSmallestPairs(int[] nums1, int[] nums2, int k) {
        List<List<Integer>> res = new ArrayList<>();
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> nums1[a[0]] + nums2[a[1]] - (nums1[b[0]] + nums2[b[1]]));
        for (int i = 0; i < Math.min(nums1.length, k); i++) {
            pq.add(new int[]{i, 0});
        }
        while (!pq.isEmpty() && k > 0) {
            int[] cur = pq.poll();
            res.add(List.of(nums1[cur[0]], nums2[cur[1]]));
            if (cur[1] + 1 < nums2.length) {
                pq.add(new int[]{cur[0], cur[1] + 1});
            }
            k--;
        }
        return res;
    }
}